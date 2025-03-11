import { describe, it, expect, vi, beforeEach } from "vitest";
import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import { MemoryRouter } from "react-router-dom";
import axios from "axios";
import SignUpModal from "../../components/auth/SignUpModal";

// Mock axios for API calls
vi.mock("axios");

// describe("SignUpModal Component", () => {
//   const onCloseMock = vi.fn();
//   const onRegisterSuccessMock = vi.fn();

//   beforeEach(() => {
//     vi.clearAllMocks();
//   });


describe("SignUpModal Component", () => {
    let onCloseMock, onRegisterSuccessMock;
  
    beforeEach(() => {
      onCloseMock = vi.fn();
      onRegisterSuccessMock = vi.fn();
      global.fetch = vi.fn(() =>
        Promise.resolve({
          ok: true,
          json: () => Promise.resolve({ message: "User registered successfully" }),
        })
      );
    });

  it("renders the signup modal correctly", () => {
    render(
      <MemoryRouter>
        <SignUpModal onClose={onCloseMock} onRegisterSuccess={onRegisterSuccessMock} />
      </MemoryRouter>
    );

    expect(screen.getByText("📝 Sign Up")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Full Name")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Email")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Password")).toBeInTheDocument();
    expect(screen.getByPlaceholderText("Contact Number")).toBeInTheDocument();
    expect(screen.getByText("Select Library")).toBeInTheDocument();
  });

  it("fetches and displays library options", async () => {
    const mockLibraries = [
      { id: 1, name: "Library One" },
      { id: 2, name: "Library Two" },
    ];
    axios.get.mockResolvedValueOnce({ data: { libraries: mockLibraries } });

    render(
      <MemoryRouter>
        <SignUpModal onClose={onCloseMock} onRegisterSuccess={onRegisterSuccessMock} />
      </MemoryRouter>
    );

    await waitFor(() => {
      expect(screen.getByText("Library One")).toBeInTheDocument();
      expect(screen.getByText("Library Two")).toBeInTheDocument();
    });
  });

  it("allows user to type in input fields", () => {
    render(
      <MemoryRouter>
        <SignUpModal onClose={onCloseMock} onRegisterSuccess={onRegisterSuccessMock} />
      </MemoryRouter>
    );

    fireEvent.change(screen.getByPlaceholderText("Full Name"), { target: { value: "John Doe" } });
    fireEvent.change(screen.getByPlaceholderText("Email"), { target: { value: "john@example.com" } });
    fireEvent.change(screen.getByPlaceholderText("Password"), { target: { value: "securepass" } });
    fireEvent.change(screen.getByPlaceholderText("Contact Number"), { target: { value: "1234567890" } });

    expect(screen.getByPlaceholderText("Full Name").value).toBe("John Doe");
    expect(screen.getByPlaceholderText("Email").value).toBe("john@example.com");
    expect(screen.getByPlaceholderText("Password").value).toBe("securepass");
    expect(screen.getByPlaceholderText("Contact Number").value).toBe("1234567890");
  });

 
  it("sends signup request and handles success", async () => {
    const mockLibraries = [{ id: 1, name: "Library One" }];
    axios.get.mockResolvedValueOnce({ data: { libraries: mockLibraries } });
  
    global.fetch = vi.fn().mockResolvedValueOnce({
      ok: true,
      json: async () => ({ message: "User registered successfully" }), // Ensure response has data
    });
  
    render(
      <MemoryRouter>
        <SignUpModal onClose={onCloseMock} onRegisterSuccess={onRegisterSuccessMock} />
      </MemoryRouter>
    );
  
    await waitFor(() => expect(screen.getByText("Library One")).toBeInTheDocument());
  
    fireEvent.change(screen.getByPlaceholderText("Full Name"), { target: { value: "John Doe" } });
    fireEvent.change(screen.getByPlaceholderText("Email"), { target: { value: "john@example.com" } });
    fireEvent.change(screen.getByPlaceholderText("Password"), { target: { value: "securepass" } });
    fireEvent.change(screen.getByPlaceholderText("Contact Number"), { target: { value: "1234567890" } });
    fireEvent.change(screen.getByRole("combobox"), { target: { value: "1" } });
  
    fireEvent.click(screen.getByText("Sign Up"));
  
    await waitFor(() => {
      expect(global.fetch).toHaveBeenCalledWith("http://localhost:8080/user/register", expect.anything());
      expect(screen.getByText("✅ Sign-up successful! Redirecting to login...")).toBeInTheDocument();
    });
  
    // Ensure onRegisterSuccessMock and onCloseMock are called
    await waitFor(() => expect(onRegisterSuccessMock).toHaveBeenCalled(), { timeout: 4000 });
    await waitFor(() => expect(onCloseMock).toHaveBeenCalled(), { timeout: 4000 });
  });
  

  it("shows error message on failed signup", async () => {
    axios.get.mockResolvedValueOnce({ data: { libraries: [{ id: 1, name: "Library One" }] } });

    global.fetch = vi.fn().mockResolvedValueOnce({
      ok: false,
      json: async () => ({ error: "Email already in use" }),
    });

    render(
      <MemoryRouter>
        <SignUpModal onClose={onCloseMock} onRegisterSuccess={onRegisterSuccessMock} />
      </MemoryRouter>
    );

    await waitFor(() => expect(screen.getByText("Library One")).toBeInTheDocument());

    fireEvent.change(screen.getByPlaceholderText("Full Name"), { target: { value: "John Doe" } });
    fireEvent.change(screen.getByPlaceholderText("Email"), { target: { value: "john@example.com" } });
    fireEvent.change(screen.getByPlaceholderText("Password"), { target: { value: "securepass" } });
    fireEvent.change(screen.getByPlaceholderText("Contact Number"), { target: { value: "1234567890" } });
    fireEvent.change(screen.getByRole("combobox"), { target: { value: "1" } });

    fireEvent.click(screen.getByText("Sign Up"));

    await waitFor(() => {
      expect(screen.getByText("Email already in use")).toBeInTheDocument();
    });
  });

  it("closes modal when clicking close button", () => {
    render(
      <MemoryRouter>
        <SignUpModal onClose={onCloseMock} onRegisterSuccess={onRegisterSuccessMock} />
      </MemoryRouter>
    );

    fireEvent.click(screen.getByText("Close"));
    expect(onCloseMock).toHaveBeenCalled();
  });
});
