"use client";
import React, { createContext, ReactNode, useEffect, useState } from "react";

interface AuthContextProps {
  email: string | null;
  username: string | null;
  authToken: string | null;
  isLoggedIn: boolean;
  Login: (identifier: string, authToken: string) => void;
  Logout: () => void;
}

export const AuthContext = createContext<AuthContextProps>({
  email: null,
  username: null,
  authToken: null,
  isLoggedIn: false,
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
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);

  function setIdentifier(identifier: string| null) {
    if (!identifier) {
      setEmail(null);
      setUsername(null);
      localStorage.removeItem("email");
      localStorage.removeItem("username");
      return
    }

    const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
    if (emailRegex.test(identifier)) {
      setEmail(identifier);
      localStorage.setItem("email", identifier);
    } else {
      setUsername(identifier);
      localStorage.setItem("username", identifier);
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
    setToken(authToken);
    setIsLoggedIn(true);
  };

  const Logout = () => {
    setIdentifier(null);
    setToken(null);
    setIsLoggedIn(false);
    localStorage.removeItem("authToken");
    localStorage.removeItem("email");
    localStorage.removeItem("username");
  }

  useEffect(() => {
    if (localStorage.getItem("authToken")) {
      setAuthToken(localStorage.getItem("authToken"));
      setIsLoggedIn(true);
    }
    if (localStorage.getItem("email")) {
      setEmail(localStorage.getItem("email"));
    }
    if (localStorage.getItem("username")) {
      setUsername(localStorage.getItem("username"));
    }
  }, []);

  return (
    <AuthContext.Provider
      value={{
        email,
        username,
        authToken,
        isLoggedIn, 
        Login,
        Logout: () => {} // Add the Logout property
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
