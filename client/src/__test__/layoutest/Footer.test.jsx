import { render, screen } from "@testing-library/react";
import Footer from "../../components/layout/Footer"; // Adjust the import path if needed

describe("Footer Component", () => {
  it("renders the footer correctly", () => {
    render(<Footer />);
    expect(screen.getByRole("contentinfo")).toBeInTheDocument();
  });

  it("displays the current year dynamically", () => {
    const currentYear = new Date().getFullYear();
    render(<Footer />);
    expect(screen.getByText(`© ${currentYear} Library Management System. All rights reserved.`)).toBeInTheDocument();
  });
});
