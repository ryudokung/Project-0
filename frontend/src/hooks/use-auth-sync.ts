"use client";

import { useEffect, useState, useCallback } from "react";
import { usePrivy } from "@privy-io/react-auth";

export function useAuthSync() {
  const { ready, authenticated, getAccessToken, user: privyUser, logout: privyLogout } = usePrivy();
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

  const logout = useCallback(() => {
    setBackendToken(null);
    setUser(null);
    localStorage.removeItem("project0_token");
    if (authenticated) {
      privyLogout();
    }
  }, [authenticated, privyLogout]);

  useEffect(() => {
    const syncWithBackend = async () => {
      if (!ready) return;

      setIsLoading(true);
      try {
        if (authenticated && privyUser) {
          const privyToken = await getAccessToken();
          const walletAddress = privyUser.wallet?.address;
          const guestId = localStorage.getItem("project0_guest_id");
          
          const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/v1/auth/login`, {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
              "Authorization": `Bearer ${privyToken}`,
            },
            body: JSON.stringify({
              privy_did: privyUser.id,
              wallet_address: walletAddress || "",
              privy_token: privyToken,
              guest_id: guestId,
            }),
          });

          if (response.ok) {
            const data = await response.json();
            setBackendToken(data.token);
            setUser(data.user);
            localStorage.setItem("project0_token", data.token);
            if (guestId) localStorage.removeItem("project0_guest_id");
          }
        } else {
          // If not authenticated via Privy, check if we have a local token
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
        }
      } catch (error) {
        console.error("Error syncing with backend:", error);
      } finally {
        setIsLoading(false);
      }
    };

    syncWithBackend();
  }, [ready, authenticated, getAccessToken, privyUser]);

  return { backendToken, user, isLoading: !ready || isLoading, guestLogin, traditionalLogin, signup, logout };
}
