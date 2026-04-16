import { render, screen, fireEvent } from "@testing-library/react";
import { describe, it, expect, vi, beforeEach } from "vitest";
import StepWizard from "./StepWizard";

describe("StepWizard Component", () => {
  // Mock data and functions
  const mockProps = {
    title: "Registration Wizard",
    stepDescriptions: {
      1: "Enter your personal details",
      2: "Choose your membership",
    },
    stepForms: {
      1: <div data-testid="form-step-1">Form Step 1 Content</div>,
      2: <div data-testid="form-step-2">Form Step 2 Content</div>,
    },
    // Mock validations: returns 0 (no errors) by default
    stepValidations: {
      1: vi.fn(() => 0),
      2: vi.fn(() => 0),
    },
    onNavigate: vi.fn(),
    onSubmit: vi.fn(),
  };

  beforeEach(() => {
    vi.clearAllMocks();
  });

  it("should render the first step correctly", () => {
    render(<StepWizard {...mockProps} />);

    expect(screen.getByTestId("title")).toHaveTextContent(
      "Registration Wizard",
    );
    expect(screen.getByText("Enter your personal details")).toBeInTheDocument();
    expect(screen.getByTestId("form-step-1")).toBeInTheDocument();
    expect(screen.getByText("1 / 2")).toBeInTheDocument();

    // Back button should not be visible on the first step
    expect(screen.queryByTestId("back-button")).not.toBeInTheDocument();
  });

  it("should navigate to the next step when validation passes", () => {
    render(<StepWizard {...mockProps} />);

    const nextButton = screen.getByTestId("next-button");
    fireEvent.click(nextButton);

    expect(mockProps.stepValidations[1]).toHaveBeenCalled();
    expect(mockProps.onNavigate).toHaveBeenCalledTimes(1);
    expect(screen.getByText("Choose your membership")).toBeInTheDocument();
    expect(screen.getByText("2 / 2")).toBeInTheDocument();
  });

  it("should stay on the current step if validation fails", () => {
    // Override validation for step 1 to return 1 (error found)
    const propsWithErrors = {
      ...mockProps,
      stepValidations: { 1: () => 1, 2: () => 0 },
    };

    render(<StepWizard {...propsWithErrors} />);

    fireEvent.click(screen.getByTestId("next-button"));

    // Validation was called, but step didn't change
    expect(screen.getByText("1 / 2")).toBeInTheDocument();
    expect(
      screen.queryByText("Choose your membership"),
    ).not.toBeInTheDocument();
    expect(mockProps.onNavigate).not.toHaveBeenCalled();
  });

  it("should allow navigating back to the previous step", () => {
    render(<StepWizard {...mockProps} />);

    // Move to step 2
    fireEvent.click(screen.getByTestId("next-button"));
    expect(screen.getByText("2 / 2")).toBeInTheDocument();

    // Click back
    const backButton = screen.getByTestId("back-button");
    fireEvent.click(backButton);

    expect(screen.getByText("1 / 2")).toBeInTheDocument();
    expect(mockProps.onNavigate).toHaveBeenCalledTimes(2);
  });

  it("should call onSubmit when clicking the submit button on the final step", () => {
    render(<StepWizard {...mockProps} />);

    // Navigate to the final step
    fireEvent.click(screen.getByTestId("next-button"));

    const submitButton = screen.getByTestId("submit-button");
    fireEvent.click(submitButton);

    expect(mockProps.onSubmit).toHaveBeenCalled();
  });

  it("should show a loading spinner and disable the button while submitting", async () => {
    render(<StepWizard {...mockProps} />);

    // Go to final step
    fireEvent.click(screen.getByTestId("next-button"));

    const submitButton = screen.getByTestId("submit-button");
    fireEvent.click(submitButton);

    // Since setIsSubmitting(true) is called immediately
    expect(submitButton).toBeDisabled();
    expect(screen.getByText(/Saving.../i)).toBeInTheDocument();
  });
});
