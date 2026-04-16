import { AlertCircleIcon } from "lucide-react";

import { Alert, AlertDescription, AlertTitle } from "../ui/alert";

interface AlertErrorProps {
  title: React.ReactNode;
  description: React.ReactNode;
}

export function AlertError({ title, description }: AlertErrorProps) {
  return (
    <Alert variant="destructive" className="max-w-md">
      <AlertCircleIcon />
      <AlertTitle>{title}</AlertTitle>
      <AlertDescription>{description}</AlertDescription>
    </Alert>
  );
}
