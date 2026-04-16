import { useState } from "react";
import { Badge } from "./ui/badge";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "./ui/card";
import { Button } from "./ui/button";
import { Check, ChevronLeft, ChevronRight } from "lucide-react";
import { Spinner } from "./ui/spinner";

interface StepWizardProps {
  title: string;
  stepDescriptions: Record<number, string>;
  stepForms: Record<number, React.ReactElement>;
  stepValidations: Record<number, () => number>;
  onNavigate: () => void;
  onSubmit: () => void;
}

function StepWizard({
  title,
  stepDescriptions,
  stepForms,
  stepValidations,
  onNavigate,
  onSubmit,
}: StepWizardProps) {
  const totalSteps = Object.keys(stepForms).length;
  const [currentStep, setCurrentStep] = useState(1);
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleNext = () => {
    const numErrors = stepValidations[currentStep]();
    if (numErrors) return;
    setCurrentStep((s) => s + 1);
    onNavigate();
  };

  const handleBack = () => {
    setCurrentStep((s) => s - 1);
    onNavigate();
  };

  const handleSubmit = async () => {
    setIsSubmitting(true);
    await onSubmit();
    setIsSubmitting(false);
  };

  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center justify-between">
          <span data-testid="title">{title}</span>
          <Badge>{`${currentStep} / ${totalSteps}`}</Badge>
        </CardTitle>
        <CardDescription>{stepDescriptions[currentStep]}</CardDescription>
      </CardHeader>
      <CardContent>{stepForms[currentStep]}</CardContent>
      <CardFooter className="gap-3 pt-4">
        {currentStep > 1 ? (
          <Button
            variant="outline"
            onClick={handleBack}
            data-testid="back-button"
            className="flex items-center gap-1.5"
          >
            <ChevronLeft />
            Back
          </Button>
        ) : (
          <div className="flex-1" />
        )}

        <div className="flex-1 flex justify-end">
          {currentStep < totalSteps ? (
            <Button
              onClick={handleNext}
              data-testid="next-button"
              className="flex items-center gap-1.5"
            >
              Next
              <ChevronRight />
            </Button>
          ) : (
            <Button
              onClick={handleSubmit}
              disabled={isSubmitting}
              data-testid="submit-button"
              className="min-w-32 flex items-center gap-2"
            >
              {isSubmitting ? (
                <>
                  <Spinner />
                  Saving...
                </>
              ) : (
                <>
                  Complete registration
                  <Check />
                </>
              )}
            </Button>
          )}
        </div>
      </CardFooter>
    </Card>
  );
}

export default StepWizard;
