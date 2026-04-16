import { CheckCircle2Icon } from "lucide-react";

import { Alert, AlertDescription, AlertTitle } from "../ui/alert";

interface AlertSuccessProps {
  title: React.ReactNode;
  description: React.ReactNode;
}

export function AlertSuccess({ title, description }: AlertSuccessProps) {
  return (
    <Alert className="max-w-md">
      <CheckCircle2Icon />
      <AlertTitle>{title}</AlertTitle>
      <AlertDescription>{description}</AlertDescription>
    </Alert>
  );
}
