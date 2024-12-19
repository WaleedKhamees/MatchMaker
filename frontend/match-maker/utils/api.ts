import { GetAuthToken } from "../src/context/Auth";

const BASE_URL = "http://localhost:8000";

export interface LoginData {
  identifier: string;
  password: string;
}

export interface RegistrationData {
  username: string;
  firstname: string;
  lastname: string;
  email: string;
  password: string;
  gender: string;
  birthdate: string;
  city: string;
  address: string;
  role: string;
}

export async function fetchMatches(): Promise<any> {
  try {
    const response = await fetch(`${BASE_URL}/match`, { cache: "no-store" });

    if (!response.ok) {
      throw new Error("Failed to fetch matches");
    }

    return await response.json();
  } catch (error) {
    console.error("Error fetching matches:", error);
    throw error;
  }
}

export async function fetchLogin(loginData: LoginData): Promise<any> {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  try {
    let data = {};
    if (emailRegex.test(loginData.identifier)) {
      data = { email: loginData.identifier, password: loginData.password };
    } else {
      data = { username: loginData.identifier, password: loginData.password };
    }

    const response = await fetch(`${BASE_URL}/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });

    if (response.ok) {
      return await response.json();
    } else {
      const error = await response.json();
      throw new Error(error.error);
    }
  } catch (error: any) {
    throw error;
  }
}

export async function fetchRegister(
  registrationData: RegistrationData
)  {
  try {
    const response = await fetch(`${BASE_URL}/register`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        ...registrationData,
        birthdate: new Date(registrationData.birthdate).toISOString(),
      }),
    });

    if (!response.ok) {
      console.log("response not ok");
      const error = await response.json();
      throw error;
    }
    const data = await response.json();
    return data;
  } catch (error) {
    throw error;
  }
}

export const fetchUnapprovedUsers = async () => {
  try {
    const response = await fetch(`http://localhost:8000/user/approve`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": GetAuthToken() || "",
        },
      }
    );
    if (!response.ok) {
      switch (response.status) {
        case 401:
          throw new Error("Unauthorized");
        case 403:
          throw new Error("Forbidden");
        default:
          throw new Error("Failed to fetch unapproved users");
      }
    }
    const data = await response.json();
    return data;
  } catch (error) {
    throw error;
  }
};
export async function fetchApproveUser(username: string, approved: boolean): Promise<any> {
  try {
    const response = await fetch(`${BASE_URL}/user/approve`,{
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        "Authorization": GetAuthToken() || "",
      },
      body: JSON.stringify({ username, approved: approved }),
    });

    if (!response.ok) {
      throw new Error("Failed to update user approval");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error updating user approval:", error);
    throw error;
  }
}