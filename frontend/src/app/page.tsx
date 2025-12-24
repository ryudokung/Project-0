"use client";

import Image from "next/image";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { LoginButton } from "@/components/login-button";
import { MechCard } from "@/components/mech-card";
import { useAuthSync } from "@/hooks/use-auth-sync";
import { usePrivy } from "@privy-io/react-auth";

export default function Home() {
  const router = useRouter();
  const { user: backendUser, isLoading: isSyncing } = useAuthSync();

  useEffect(() => {
    if (!isSyncing && backendUser) {
      if (!backendUser.active_character_id) {
        router.push('/character-creation');
      } else {
        router.push('/hangar');
      }
    }
  }, [isSyncing, backendUser, router]);

  return (
    <div className="flex min-h-screen flex-col items-center font-sans text-white">
      <header className="w-full max-w-7xl flex justify-between items-center p-6">
        <div className="flex items-center gap-2">
          <div className="relative h-8 w-8">
            <Image className="dark:invert" src="/next.svg" alt="Logo" fill />
          </div>
          <span className="font-bold tracking-tighter text-xl">PROJECT-0</span>
        </div>
        <LoginButton />
      </header>

      <main className="flex flex-col items-center justify-center w-full max-w-7xl px-6 flex-1">
        <div className="flex flex-col items-center gap-8 text-center mb-20">
          <h1 className="text-5xl font-bold tracking-tighter sm:text-7xl uppercase italic">
            PROJECT-<span className="text-pink-500">0</span>
          </h1>
          <p className="max-w-[600px] text-zinc-400 md:text-xl font-light tracking-wide">
            The next generation of AI-driven, Web3-powered space exploration and combat. 
            Create your pilot, command your fleet, and conquer the void.
          </p>
          <div className="flex gap-4">
            <LoginButton />
          </div>
        </div>

        {/* Decorative Elements */}
        <div className="absolute bottom-0 left-0 w-full h-64 bg-gradient-to-t from-pink-500/10 to-transparent pointer-events-none" />
      </main>
      
      <footer className="py-10 text-sm text-zinc-600">
        <p>© 2024 Project-0. All rights reserved.</p>
      </footer>
    </div>
  );
}
