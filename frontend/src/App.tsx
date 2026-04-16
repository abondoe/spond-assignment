import { useEffect, useState } from "react";
import "./App.css";
import StepWizard from "./components/StepWizard";
import type {
  CreateRegistrationRequest,
  GetFormResponse,
} from "./generated-types/api";
import MemberTypeSelect from "./components/MemberTypeSelect";
import UserInfoForm, {
  type UserData,
  type UserInfoErrors,
} from "./components/UserInfoForm";
import Summary from "./components/Summary";
import { AlertError } from "./components/alert/AlertError";
import { AlertSuccess } from "./components/alert/AlertSuccess";
import { AlertInformation } from "./components/alert/AlertInformation";

function AppContainer({ children }: { children: React.ReactNode }) {
  return (
    <div className="min-h-screen w-full flex justify-center items-start bg-slate-50 p-4 pt-20">
      <div className="w-full max-w-md">{children}</div>
    </div>
  );
}

function App() {
  const formId = window.location.pathname.split("/")[1];
  const today = new Date();
  today.setHours(0, 0, 0, 0);
  const [form, setForm] = useState<GetFormResponse | null>(null);
  const [loading, setLoading] = useState(true);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [registerSuccess, setRegisterSuccess] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [selectedMemberTypeId, setSelectedMemberTypeId] = useState<
    string | null
  >(null);
  const [userInfo, setUserInfo] = useState<UserData>({
    name: "",
    email: "",
    phoneNumber: "",
    birthDate: "",
  });
  const [userInfoErrors, setUserInfoErrors] = useState<UserInfoErrors>({});
  const [memberTypeError, setMemberTypeError] = useState("");

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        const response = await fetch("/api/forms/" + formId);

        if (!response.ok) {
          if (response.status == 400) {
            throw new Error(
              "The registration form you are looking for does not exist",
            );
          } else {
            throw new Error(`HTTP error! status: ${response.status}`);
          }
        }

        const data = await response.json();
        setForm(data);
      } catch (err) {
        if (err instanceof Error) {
          setError(err.message);
        } else {
          setError("An unknown error occured");
        }
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [formId]);

  const submitRegistration = async () => {
    if (!selectedMemberTypeId || !form?.formId) {
      setError("Missing member type or form information.");
      return;
    }

    try {
      setIsSubmitting(true);
      setError(null);

      const payload: CreateRegistrationRequest = {
        formId: form.formId,
        memberTypeId: selectedMemberTypeId,
        name: userInfo.name,
        email: userInfo.email,
        phoneNumber: userInfo.phoneNumber,
        birthDate: new Date(userInfo.birthDate).toISOString(), // Ensures RFC3339 format
      };

      const response = await fetch("/api/registrations", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(payload),
      });

      if (!response.ok) {
        if (response.status == 409) {
          throw new Error(
            "Multiple registrations of the same person is not allowed",
          );
        }
        const errorData = await response.json();
        throw new Error(errorData.message || "Failed to save registration");
      }

      setRegisterSuccess(true);
    } catch (err) {
      setError(err instanceof Error ? err.message : "An error occurred");
    } finally {
      setIsSubmitting(false);
    }
  };

  const stepDescriptions = {
    1: "Select you preferred member type to start registration",
    2: "Tell us who you are by filling in your details",
    3: "Review your information before completing the registration",
  };

  const validateMemberType = () => {
    if (!selectedMemberTypeId) {
      setMemberTypeError("Membertype is required");
      return 1;
    }
    return 0;
  };

  const validateEmail = (email: string) => {
    const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return re.test(email.toLowerCase());
  };

  const validateUserInfo = () => {
    const newErrors: UserInfoErrors = {};

    if (!userInfo.name) newErrors.name = "Name is required";
    if (!userInfo.email) {
      newErrors.email = "Email is required";
    } else if (!validateEmail(userInfo.email)) {
      newErrors.email = "Please provide a valid email address";
    }
    if (userInfo.phoneNumber.length < 8)
      newErrors.phoneNumber = "Minimum 8 digits required";
    if (!userInfo.birthDate) {
      newErrors.birthDate = "Birth date is required";
    } else {
      const selectedDate = new Date(userInfo.birthDate);
      if (selectedDate >= today) newErrors.birthDate = "Must be in the past";
    }

    setUserInfoErrors(newErrors);
    return Object.keys(newErrors).length;
  };

  const stepValidations = {
    1: validateMemberType,
    2: validateUserInfo,
    3: () => 0,
  };

  const selectedMembertypeName = form?.memberTypes.find(
    (memberType) => memberType.id == selectedMemberTypeId,
  )?.name;

  const registrationOpensDate = new Date(form?.registrationOpens || "");

  const stepForms = {
    1: (
      <MemberTypeSelect
        memberTypes={form?.memberTypes}
        value={selectedMemberTypeId}
        onChange={setSelectedMemberTypeId}
        error={memberTypeError}
      />
    ),
    2: (
      <UserInfoForm
        data={userInfo}
        onChange={setUserInfo}
        errors={userInfoErrors}
      />
    ),
    3: <Summary memberType={selectedMembertypeName} userData={userInfo} />,
  };

  function resetErrors() {
    setUserInfoErrors({});
    setMemberTypeError("");
  }

  if (registrationOpensDate > today)
    return (
      <AppContainer>
        <AlertInformation
          title="Information"
          description={
            <span>
              The registration for{" "}
              <strong className="font-bold">{form?.title}</strong> is not open
              yet. Registration opens on{" "}
              <strong className="font-bold">
                {registrationOpensDate.toLocaleDateString()}
              </strong>
            </span>
          }
        />
      </AppContainer>
    );
  if (loading)
    return (
      <AppContainer>
        <p>Laster form...</p>
      </AppContainer>
    );
  if (error)
    return (
      <AppContainer>
        <AlertError title="Error" description={error} />
      </AppContainer>
    );
  if (!isSubmitting && registerSuccess)
    return (
      <AppContainer>
        <AlertSuccess
          title="Registration successful"
          description={`You have successfully registered for ${form?.title}`}
        />
      </AppContainer>
    );
  return (
    <AppContainer>
      {form && (
        <StepWizard
          title={form.title}
          stepDescriptions={stepDescriptions}
          stepForms={stepForms}
          stepValidations={stepValidations}
          onNavigate={resetErrors}
          onSubmit={submitRegistration}
        />
      )}
    </AppContainer>
  );
}

export default App;
