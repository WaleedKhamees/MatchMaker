"use client";
import React, { createContext, ReactNode, useState } from "react";

interface AuthContextProps {
  email: string | null;
  username: string | null;
  authToken: string | null;
  Login: (identifier: string, authToken: string) => void;
  Logout: () => void;
}

export const AuthContext = createContext<AuthContextProps>({
  email: null,
  username: null,
  authToken: null,
  Login: () => {},
  Logout: () => {},
});

export const AuthProvider = ({
  children,
}: {
  children: ReactNode | ReactNode[];
}) => {
  const [email, setEmail] = useState<string | null>(null);
  const [username, setUsername] = useState<string | null>(null);
  const [authToken, setAuthToken] = useState<string | null>(null);

  function setIdentifier(identifier: string) {
    const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
    if (emailRegex.test(identifier)) {
      setEmail(identifier);
    } else {
      setUsername(identifier);
    }
  }

  function setToken(authToken: string | null) {
    const localStorageAuth = localStorage.getItem("authToken");
    if (localStorageAuth) {
      setAuthToken(localStorageAuth);
    }
    setAuthToken(authToken);
    if (authToken) {
      setAuthToken(authToken);
      localStorage.setItem("authToken", authToken);
    }
  }

  const Login = async (identifier: string, authToken: string) => {
    setIdentifier(identifier);
    setAuthToken(authToken);
  };

  return (
    <AuthContext.Provider
      value={{
        email,
        username,
        authToken,
        Login,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};

export const GetAuthToken = () => {
  const authToken = localStorage.getItem("authToken");
  return authToken;
};
