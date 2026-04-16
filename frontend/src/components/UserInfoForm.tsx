import { Field, FieldDescription, FieldLabel } from "./ui/field";
import { Input } from "./ui/input";

export interface UserData {
  name: string;
  email: string;
  phoneNumber: string;
  birthDate: string;
}

export type UserInfoErrors = Partial<Record<keyof UserData, string>>;

interface UserInfoFormProps {
  data: UserData;
  onChange: (data: UserData) => void;
  errors?: UserInfoErrors;
}

function UserInfoForm({ data, onChange, errors }: UserInfoFormProps) {
  const handleChange = (field: keyof UserData, value: string) => {
    onChange({ ...data, [field]: value });
  };

  return (
    <div className="space-y-4">
      <Field>
        <FieldLabel htmlFor="input-name">Name</FieldLabel>
        <Input
          id="input-name"
          value={data.name}
          onChange={(e) => handleChange("name", e.target.value)}
          placeholder="Luke Skywalker"
          aria-invalid={!!errors?.name}
        />
        {errors?.name && (
          <FieldDescription className="text-destructive">
            {errors.name}
          </FieldDescription>
        )}
      </Field>

      <Field>
        <FieldLabel htmlFor="input-email">Email</FieldLabel>
        <Input
          id="input-email"
          type="email"
          value={data.email}
          onChange={(e) => handleChange("email", e.target.value)}
          placeholder="luke@example.com"
          aria-invalid={!!errors?.email}
        />
        {errors?.email && (
          <FieldDescription className="text-destructive">
            {errors.email}
          </FieldDescription>
        )}
      </Field>

      <Field>
        <FieldLabel htmlFor="input-phone-number">Phone number</FieldLabel>
        <Input
          id="input-phone-number"
          type="tel"
          value={data.phoneNumber}
          onChange={(e) => handleChange("phoneNumber", e.target.value)}
          placeholder="12345678"
          aria-invalid={!!errors?.phoneNumber}
        />
        {errors?.phoneNumber && (
          <FieldDescription className="text-destructive">
            {errors.phoneNumber}
          </FieldDescription>
        )}
      </Field>

      <Field>
        <FieldLabel htmlFor="input-birth-date">Birth date</FieldLabel>
        <Input
          id="input-birth-date"
          type="date"
          value={data.birthDate}
          onChange={(e) => handleChange("birthDate", e.target.value)}
          aria-invalid={!!errors?.birthDate}
        />
        {errors?.birthDate && (
          <FieldDescription className="text-destructive">
            {errors.birthDate}
          </FieldDescription>
        )}
      </Field>
    </div>
  );
}

export default UserInfoForm;
