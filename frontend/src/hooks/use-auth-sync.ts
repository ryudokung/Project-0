"use client";

import { useEffect, useState } from "react";
import { usePrivy } from "@privy-io/react-auth";

export function useAuthSync() {
  const { ready, authenticated, getAccessToken, user: privyUser } = usePrivy();
  const [backendToken, setBackendToken] = useState<string | null>(null);
  const [user, setUser] = useState<any | null>(null);
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    const syncWithBackend = async () => {
      if (ready && authenticated && privyUser) {
        setIsLoading(true);
        try {
          const privyToken = await getAccessToken();
          const walletAddress = privyUser.wallet?.address;
          
          const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/v1/auth/login`, {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
              "Authorization": `Bearer ${privyToken}`,
            },
            body: JSON.stringify({
              wallet_address: walletAddress,
              privy_token: privyToken,
            }),
          });

          if (response.ok) {
            const data = await response.json();
            setBackendToken(data.token);
            setUser(data.user);
            localStorage.setItem("project0_token", data.token);
          } else {
            console.error("Failed to sync with backend");
          }
        } catch (error) {
          console.error("Error syncing with backend:", error);
        } finally {
          setIsLoading(false);
        }
      } else if (ready && !authenticated) {
        setBackendToken(null);
        setUser(null);
        localStorage.removeItem("project0_token");
      }
    };

    syncWithBackend();
  }, [ready, authenticated, getAccessToken, privyUser]);

  return { backendToken, user, isLoading };
}
