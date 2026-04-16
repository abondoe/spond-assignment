import { test, expect } from "@playwright/test";

test.beforeEach(async ({ request }) => {
  const response = await request.post("http://localhost:8082/api/test/reset");
  expect(response.ok()).toBeTruthy();
});

test("a form that opens in the future", async ({ page }) => {
  // One day before the registration form becomes open
  await page.clock.setFixedTime(new Date("2024-12-15T00:00:00"));
  await page.goto("http://localhost:3002/B171388180BC457D9887AD92B6CCFC86");
  await test.step("page should show error banner to user", async () => {
    await expect(page.getByRole("alert")).toContainText("Information");
    await expect(page.getByRole("alert")).toContainText(
      "The registration for Coding camp summer 2025 is not open yet. Registration opens on 12/16/2024",
    );
  });
});

test("a non existing form registration", async ({ page }) => {
  await page.goto("http://localhost:3002/B171388180BC457D9887AD92B6CCFC8");
  await test.step("page should show error banner to user", async () => {
    await expect(page.getByRole("alert")).toContainText("Error");
    await expect(page.getByRole("alert")).toContainText(
      "The registration form you are looking for does not exist",
    );
  });
});

test("an existing form registration - wizard should take the user step wise through registration", async ({
  page,
}) => {
  await page.goto("http://localhost:3002/B171388180BC457D9887AD92B6CCFC86");

  await test.step("wizard step 1.1: should have titles, labels and button", async () => {
    await expect(page.getByTestId("title")).toContainText(
      "Coding camp summer 2025",
    );
    await expect(page.locator("#root")).toContainText("1 / 3");
    await expect(page.locator("#root")).toContainText(
      "Select you preferred member type to start registration",
    );
    await expect(page.locator("label")).toContainText("Member type");
    await expect(page.getByTestId("next-button")).toBeVisible();
  });
  await test.step("wizard step 1.2: should have member types", async () => {
    await page.getByRole("combobox", { name: "Member type" }).click();
    await expect(page.locator("#base-ui-_r_0_-list")).toContainText(
      "Active Member",
    );
    await expect(page.locator("#base-ui-_r_0_-list")).toContainText(
      "Social Member",
    );
    await page.locator('[id="_r_3_"] > div').first().click();
  });
  await test.step("wizard step 1.3: should require member type", async () => {
    await page.getByTestId("next-button").click();
    await expect(page.getByRole("paragraph")).toContainText(
      "Membertype is required",
    );
  });
  await test.step("wizard step 1.4: should move to next step", async () => {
    await page.getByRole("combobox", { name: "Member type" }).click();
    await page.getByText("Active Member").click();
    await page.getByTestId("next-button").click();
    await expect(page.getByTestId("back-button")).toBeVisible();
  });
  await test.step("wizard step 2.1: should have titles, labels and buttons", async () => {
    await expect(page.getByTestId("title")).toContainText(
      "Coding camp summer 2025",
    );
    await expect(page.locator("#root")).toContainText("2 / 3");
    await expect(page.locator("#root")).toContainText(
      "Tell us who you are by filling in your details",
    );
    await expect(page.locator("#root")).toContainText("Name");
    await expect(page.locator("#root")).toContainText("Email");
    await expect(page.locator("#root")).toContainText("Phone number");
    await expect(page.locator("#root")).toContainText("Birth date");
    await expect(page.getByRole("textbox", { name: "Name" })).toBeVisible();
    await expect(page.getByRole("textbox", { name: "Email" })).toBeVisible();
    await expect(
      page.getByRole("textbox", { name: "Phone number" }),
    ).toBeVisible();
    await expect(
      page.getByRole("textbox", { name: "Birth date" }),
    ).toBeVisible();
    await expect(page.getByTestId("back-button")).toBeVisible();
    await expect(page.getByTestId("next-button")).toBeVisible();
  });
  await test.step("step2.2: should allow back navigation", async () => {
    await page.getByTestId("back-button").click();
    await expect(page.locator("#root")).toContainText("1 / 3");
    await page.getByTestId("next-button").click();
    await expect(page.locator("#root")).toContainText("2 / 3");
  });
  await test.step("wizard step wizard 2.3: should require all fields", async () => {
    await page.getByTestId("next-button").click();
    await expect(page.locator("#root")).toContainText("Name is required");
    await expect(page.locator("#root")).toContainText("Email is required");
    await expect(page.locator("#root")).toContainText(
      "Minimum 8 digits required",
    );
    await expect(page.locator("#root")).toContainText("Birth date is required");
  });
  await test.step("wizard step 2.4: should move to next step", async () => {
    await page.getByRole("textbox", { name: "Name" }).fill("Jane Davids");
    await page.getByRole("textbox", { name: "Email" }).fill("jane@example.com");
    await page.getByRole("textbox", { name: "Phone number" }).fill("12345678");
    await page.getByRole("textbox", { name: "Birth date" }).fill("1990-01-01");
    await page.getByTestId("next-button").click();
    await expect(page.getByTestId("submit-button")).toBeVisible();
  });
  await test.step("wizard step 3.1: should have titles, labels, summary and buttons", async () => {
    await expect(page.getByTestId("title")).toContainText(
      "Coding camp summer 2025",
    );
    await expect(page.locator("#root")).toContainText("3 / 3");
    await expect(page.locator("#root")).toContainText(
      "Review your information before completing the registration",
    );
    await expect(page.locator("dl")).toContainText("Member type");
    await expect(page.locator("dl")).toContainText("Active Member");
    await expect(page.locator("dl")).toContainText("Name");
    await expect(page.locator("dl")).toContainText("Jane Davids");
    await expect(page.locator("dl")).toContainText("Email");
    await expect(page.locator("dl")).toContainText("jane@example.com");
    await expect(page.locator("dl")).toContainText("Phone number");
    await expect(page.locator("dl")).toContainText("12345678");
    await expect(page.locator("dl")).toContainText("Birth date");
    await expect(page.locator("dl")).toContainText("1/1/1990");
    await expect(page.getByTestId("back-button")).toBeVisible();
    await expect(page.getByTestId("submit-button")).toBeVisible();
  });
  await test.step("wizard step 3.2: should allow back navigation", async () => {
    await page.getByTestId("back-button").click();
    await expect(page.locator("#root")).toContainText("2 / 3");
    await page.getByTestId("next-button").click();
    await expect(page.locator("#root")).toContainText("3 / 3");
  });
  await test.step("wizard step 3.3: should complete registration", async () => {
    await page.getByTestId("submit-button").click();
    await expect(page.getByRole("alert")).toContainText(
      "Registration successful",
    );
    await expect(page.getByRole("alert")).toContainText(
      "You have successfully registered for Coding camp summer 2025",
    );
  });
  await test.step("page should not allow duplicate registration", async () => {
    await page.goto("http://localhost:3002/B171388180BC457D9887AD92B6CCFC86");
    await page.getByRole("combobox", { name: "Member type" }).click();
    await page.getByText("Active Member").click();
    await page.getByTestId("next-button").click();
    await page.getByRole("textbox", { name: "Name" }).fill("Jane Davids");
    await page.getByRole("textbox", { name: "Email" }).fill("jane@example.com");
    await page.getByRole("textbox", { name: "Phone number" }).fill("12345678");
    await page.getByRole("textbox", { name: "Birth date" }).fill("1990-01-01");
    await page.getByTestId("next-button").click();
    await page.getByTestId("submit-button").click();
    await expect(page.getByRole("alert")).toContainText("Error");
    await expect(page.getByRole("alert")).toContainText(
      "Multiple registrations of the same person is not allowed",
    );
  });
});
