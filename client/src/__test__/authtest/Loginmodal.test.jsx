import { describe, it, expect, vi, beforeEach } from "vitest";
import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import { MemoryRouter } from "react-router-dom";
import LoginModal from "../../components/auth/LoginModal";

// Mock `jwt-decode`
vi.mock("jwt-decode", () => ({
  jwtDecode: vi.fn(() => ({ role: "Reader" })), // Default mock role
}));

// Mock `useNavigate`
const mockNavigate = vi.fn();
vi.mock("react-router-dom", async () => {
  const actual = await vi.importActual("react-router-dom");
  return {
    ...actual,
    useNavigate: () => mockNavigate,
  };
});

// Mock fetch
global.fetch = vi.fn();

describe("LoginModal Component", () => {
  const onCloseMock = vi.fn();
  const onLoginMock = vi.fn();

  beforeEach(() => {
    vi.clearAllMocks();
  });

  it("renders login modal correctly", () => {
    render(
      <MemoryRouter>
        <LoginModal onClose={onCloseMock} onLogin={onLoginMock} />
      </MemoryRouter>
    );

    expect(screen.getByText("🔐 Login")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Enter your email")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Enter your password")).toBeInTheDocument();
  });

  it("allows user to type email and password", () => {
    render(
      <MemoryRouter>
        <LoginModal onClose={onCloseMock} onLogin={onLoginMock} />
      </MemoryRouter>
    );

    const emailInput = screen.getByPlaceholderText("Enter your email");
    const passwordInput = screen.getByPlaceholderText("Enter your password");

    fireEvent.change(emailInput, { target: { value: "test@example.com" } });
    fireEvent.change(passwordInput, { target: { value: "password123" } });

    expect(emailInput.value).toBe("test@example.com");
    expect(passwordInput.value).toBe("password123");
  });

  it("calls API on login attempt with correct credentials", async () => {
    fetch.mockResolvedValueOnce({
      ok: true,
      json: async () => ({ token: "mockToken" }),
    });

    render(
      <MemoryRouter>
        <LoginModal onClose={onCloseMock} onLogin={onLoginMock} />
      </MemoryRouter>
    );

    fireEvent.change(screen.getByPlaceholderText("Enter your email"), {
      target: { value: "test@example.com" },
    });
    fireEvent.change(screen.getByPlaceholderText("Enter your password"), {
      target: { value: "password123" },
    });

    fireEvent.click(screen.getByText("Login"));

    await waitFor(() => {
      expect(fetch).toHaveBeenCalledWith("http://localhost:8080/auth/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email: "test@example.com", password: "password123" }),
      });
    });

    expect(localStorage.getItem("token")).toBe("mockToken");
    expect(onLoginMock).toHaveBeenCalled();
    expect(onCloseMock).toHaveBeenCalled();
  });

  it("displays an error message on failed login", async () => {
    fetch.mockResolvedValueOnce({
      ok: false,
      json: async () => ({ error: "Invalid credentials" }),
    });

    render(
      <MemoryRouter>
        <LoginModal onClose={onCloseMock} onLogin={onLoginMock} />
      </MemoryRouter>
    );

    fireEvent.change(screen.getByPlaceholderText("Enter your email"), {
      target: { value: "wrong@example.com" },
    });
    fireEvent.change(screen.getByPlaceholderText("Enter your password"), {
      target: { value: "wrongpassword" },
    });

    fireEvent.click(screen.getByText("Login"));

    await waitFor(() => {
      expect(screen.getByText("Invalid credentials")).toBeInTheDocument();
    });
  });

  it("navigates based on user role after login", async () => {
    fetch.mockResolvedValueOnce({
      ok: true,
      json: async () => ({ token: "mockToken" }),
    });

    const { jwtDecode } = await import("jwt-decode");
    jwtDecode.mockReturnValueOnce({ role: "LibraryAdmin" });

    render(
      <MemoryRouter>
        <LoginModal onClose={onCloseMock} onLogin={onLoginMock} />
      </MemoryRouter>
    );

    fireEvent.change(screen.getByPlaceholderText("Enter your email"), {
      target: { value: "admin@example.com" },
    });
    fireEvent.change(screen.getByPlaceholderText("Enter your password"), {
      target: { value: "password123" },
    });

    fireEvent.click(screen.getByText("Login"));

    await waitFor(() => {
      expect(mockNavigate).toHaveBeenCalledWith("/admindashboard");
    });
  });

  it("closes modal when clicking close button", () => {
    render(
      <MemoryRouter>
        <LoginModal onClose={onCloseMock} onLogin={onLoginMock} />
      </MemoryRouter>
    );

    fireEvent.click(screen.getByText("Close"));
    expect(onCloseMock).toHaveBeenCalled();
  });
});
