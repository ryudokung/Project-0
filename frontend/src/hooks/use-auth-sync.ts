"use client";

import { useEffect, useState, useCallback } from "react";

export function useAuthSync() {
  const [backendToken, setBackendToken] = useState<string | null>(null);
  const [user, setUser] = useState<any | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  // Guest Login Function
  const guestLogin = useCallback(async () => {
    let guestId = localStorage.getItem("project0_guest_id");
    if (!guestId) {
      guestId = `guest_${Math.random().toString(36).substring(2, 15)}`;
      localStorage.setItem("project0_guest_id", guestId);
    }

    setIsLoading(true);
    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/v1/auth/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ guest_id: guestId }),
      });

      if (response.ok) {
        const data = await response.json();
        setBackendToken(data.token);
        setUser(data.user);
        localStorage.setItem("project0_token", data.token);
      }
    } catch (error) {
      console.error("Guest login failed:", error);
    } finally {
      setIsLoading(false);
    }
  }, []);

  // Traditional Login Function
  const traditionalLogin = useCallback(async (username: string, password: string) => {
    setIsLoading(true);
    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/v1/auth/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password }),
      });

      if (response.ok) {
        const data = await response.json();
        setBackendToken(data.token);
        setUser(data.user);
        localStorage.setItem("project0_token", data.token);
        return { success: true };
      } else {
        const err = await response.text();
        return { success: false, error: err };
      }
    } catch (error: any) {
      return { success: false, error: error.message };
    } finally {
      setIsLoading(false);
    }
  }, []);

  // Signup Function (with Guest Upgrade support)
  const signup = useCallback(async (username: string, email: string, password: string) => {
    const guestId = localStorage.getItem("project0_guest_id");
    setIsLoading(true);
    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/v1/auth/signup`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, email, password, guest_id: guestId }),
      });

      if (response.ok) {
        const data = await response.json();
        setBackendToken(data.token);
        setUser(data.user);
        localStorage.setItem("project0_token", data.token);
        // Clear guest ID after successful upgrade
        localStorage.removeItem("project0_guest_id");
        return { success: true };
      } else {
        const err = await response.text();
        return { success: false, error: err };
      }
    } catch (error: any) {
      return { success: false, error: error.message };
    } finally {
      setIsLoading(false);
    }
  }, []);

  const refreshUser = useCallback(async () => {
    const localToken = localStorage.getItem("project0_token");
    if (localToken) {
      try {
        const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/v1/auth/me`, {
          headers: {
            "Authorization": `Bearer ${localToken}`,
          },
        });
        if (response.ok) {
          const userData = await response.json();
          setUser(userData);
        }
      } catch (error) {
        console.error("Error refreshing user:", error);
      }
    }
  }, []);

  const logout = useCallback(() => {
    setBackendToken(null);
    setUser(null);
    localStorage.removeItem("project0_token");
  }, []);

  useEffect(() => {
    const syncWithBackend = async () => {
      setIsLoading(true);
      try {
        // Check if we have a local token
        const localToken = localStorage.getItem("project0_token");
        if (localToken) {
          const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/v1/auth/me`, {
            headers: {
              "Authorization": `Bearer ${localToken}`,
            },
          });
          if (response.ok) {
            const userData = await response.json();
            setBackendToken(localToken);
            setUser(userData);
          } else {
            localStorage.removeItem("project0_token");
            setBackendToken(null);
            setUser(null);
          }
        } else {
          setBackendToken(null);
          setUser(null);
        }
      } catch (error) {
        console.error("Error syncing with backend:", error);
      } finally {
        setIsLoading(false);
      }
    };

    syncWithBackend();
  }, []);

  return { backendToken, user, isLoading, guestLogin, traditionalLogin, signup, logout, refreshUser };
}
