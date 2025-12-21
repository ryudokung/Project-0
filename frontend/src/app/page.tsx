"use client";

import Image from "next/image";
import { LoginButton } from "@/components/login-button";
import { useAuthSync } from "@/hooks/use-auth-sync";
import { usePrivy } from "@privy-io/react-auth";

export default function Home() {
  const { authenticated, user } = usePrivy();
  const { backendToken, isLoading } = useAuthSync();

  return (
    <div className="flex min-h-screen flex-col items-center justify-center bg-zinc-50 font-sans dark:bg-black">
      <main className="flex flex-col items-center gap-8 text-center">
        <div className="relative h-32 w-32">
          <Image
            className="dark:invert"
            src="/next.svg"
            alt="Project-0 Logo"
            fill
            priority
          />
        </div>
        
        <div className="flex flex-col gap-4">
          <h1 className="text-4xl font-bold tracking-tighter sm:text-6xl text-black dark:text-zinc-50">
            PROJECT-0
          </h1>
          <p className="max-w-[600px] text-zinc-600 dark:text-zinc-400 md:text-xl/relaxed lg:text-base/relaxed xl:text-xl/relaxed">
            The next generation of AI-driven, Web3-powered space exploration and combat.
          </p>
        </div>

        <div className="flex flex-col gap-4 items-center">
          <LoginButton />
          
          {authenticated && user && (
            <div className="mt-4 p-4 border border-zinc-200 dark:border-zinc-800 rounded-lg bg-white dark:bg-zinc-900">
              <p className="text-sm text-zinc-500">Logged in as:</p>
              <p className="font-mono text-xs">{user.wallet?.address}</p>
              {isLoading ? (
                <p className="text-xs text-yellow-500 mt-2">Syncing with backend...</p>
              ) : backendToken ? (
                <p className="text-xs text-green-500 mt-2">✓ Backend Authenticated</p>
              ) : (
                <p className="text-xs text-red-500 mt-2">✗ Backend Sync Failed</p>
              )}
            </div>
          )}
        </div>
      </main>
      
      <footer className="mt-20 flex gap-6 text-sm text-zinc-500">
        <p>© 2024 Project-0. All rights reserved.</p>
      </footer>
    </div>
  );
}
