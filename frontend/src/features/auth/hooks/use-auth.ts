"use client";

import { useState, useCallback } from "react";
import type { AuthUser } from "@/features/auth/types";

export function useAuth() {
  const [user, setUser] = useState<AuthUser | null>(null);
  const [isLoading, setIsLoading] = useState(false);

  const login = useCallback(async (email: string, password: string) => {
    setIsLoading(true);
    try {
      // TODO: Implement login API call to Go backend
      console.log("login", email, password);
    } finally {
      setIsLoading(false);
    }
  }, []);

  const logout = useCallback(async () => {
    setUser(null);
    // TODO: Implement logout API call
  }, []);

  return { user, isLoading, login, logout };
}
