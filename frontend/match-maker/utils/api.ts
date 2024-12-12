const BASE_URL = 'http://localhost:8000';

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
    const response = await fetch(`${BASE_URL}/match`, { cache: 'no-store' });

    if (!response.ok) {
      throw new Error('Failed to fetch matches');
    }

    return await response.json();
  } catch (error) {
    console.error('Error fetching matches:', error);
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
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });

    if (response.ok) {
      return await response.json();
    } else {
      const errorText = await response.text();
      throw new Error(errorText || 'Login failed');
    }
  } catch (error) {
    console.error('Login error:', error);
    throw error;
  }
}

export async function register(registrationData: RegistrationData): Promise<void> {
  try {
    const response = await fetch(`${BASE_URL}/register`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        ...registrationData,
        birthdate: new Date(registrationData.birthdate).toISOString(),
      }),
    });

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(errorText || 'Registration failed');
    }
  } catch (error) {
    console.error('Registration error:', error);
    throw error;
  }
}
