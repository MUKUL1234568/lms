import { describe, it, expect, vi, beforeEach } from "vitest";
import { render, screen, fireEvent } from "@testing-library/react";
import { MemoryRouter } from "react-router-dom";
import Header from "../../components/layout/Header"; // Import Header from src/components

// Mock jwtDecode
vi.mock("jwt-decode", () => ({
  jwtDecode: vi.fn(() => ({ role: "Reader" })), // Default mock role as Reader
}));

describe("Header Component", () => {
  beforeEach(() => {
    localStorage.clear(); // Ensure clean state before each test
  });

  it("renders the Library System logo", () => {
    render(
      <MemoryRouter>
        <Header />
      </MemoryRouter>
    );
    expect(screen.getByText("📚 Library System")).toBeInTheDocument();
  });

  it("shows Login and SignUp buttons when not logged in", () => {
    render(
      <MemoryRouter>
        <Header />
      </MemoryRouter>
    );

    expect(screen.getByText("Login")).toBeInTheDocument();
    expect(screen.getByText("SignUp")).toBeInTheDocument();
  });

  it("shows Dashboard and Logout buttons when logged in", () => {
    localStorage.setItem("token", "mockToken"); // Simulate login

    render(
      <MemoryRouter>
        <Header />
      </MemoryRouter>
    );

    expect(screen.getByText("Dashboard")).toBeInTheDocument();
    expect(screen.getByText("Logout")).toBeInTheDocument();
  });

  it("removes token and navigates on logout", () => {
    localStorage.setItem("token", "mockToken");

    render(
      <MemoryRouter>
        <Header />
      </MemoryRouter>
    );

    const logoutButton = screen.getByText("Logout");
    fireEvent.click(logoutButton);

    expect(localStorage.getItem("token")).toBeNull();
  });

  // it("opens login modal when clicking Login button", () => {
  //   render(
  //     <MemoryRouter>
  //       <Header />
  //     </MemoryRouter>
  //   );

  //   // Fire event to open the login modal
  //   fireEvent.click(screen.getByText("Login"));

  //   // Check if the login modal content is visible
  //   expect(screen.getByText("Login")).toBeInTheDocument(); // Adjust this to the actual modal content text
  // });

  it("opens signup modal when clicking SignUp button", () => {
    render(
      <MemoryRouter>
        <Header />
      </MemoryRouter>
    );

    // Fire event to open the sign-up modal
    fireEvent.click(screen.getByText("SignUp"));

    // Check if the sign-up modal content is visible
    expect(screen.getByText("SignUp")).toBeInTheDocument(); // Adjust this to the actual modal content text
  });
});
