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
