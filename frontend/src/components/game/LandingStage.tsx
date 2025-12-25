'use client';

import Image from "next/image";
import { LoginButton } from "@/components/login-button";
import { useAuthSync } from "@/hooks/use-auth-sync";
import { usePrivy } from "@privy-io/react-auth";

interface LandingStageProps {
  onStart: () => void;
}

export default function LandingStage({ onStart }: LandingStageProps) {
  const { login: privyLogin, authenticated: privyAuthenticated } = usePrivy();
  const { user, guestLogin, isLoading } = useAuthSync();

  const handleStart = () => {
    if (user) {
      onStart();
    } else {
      guestLogin();
    }
  };

  const handleNeuralLink = () => {
    if (user) {
      onStart();
    } else {
      privyLogin();
    }
  };

  return (
    <div className="flex min-h-screen flex-col items-center font-sans text-white bg-black">
      <header className="w-full max-w-7xl flex justify-between items-center p-6 z-20">
        <div className="flex items-center gap-2">
          <div className="relative h-8 w-8">
            <Image className="dark:invert" src="/next.svg" alt="Logo" fill />
          </div>
          <span className="font-bold tracking-tighter text-xl">PROJECT-0</span>
        </div>
        <LoginButton />
      </header>

      <main className="flex flex-col items-center justify-center w-full max-w-7xl px-6 flex-1 z-10">
        <div className="flex flex-col items-center gap-8 text-center mb-20">
          <h1 className="text-5xl font-bold tracking-tighter sm:text-7xl uppercase italic">
            PROJECT-<span className="text-pink-500">0</span>
          </h1>
          <p className="max-w-[600px] text-zinc-400 md:text-xl font-light tracking-wide">
            The next generation of AI-driven, Web3-powered space exploration and combat. 
            Create your pilot, command your fleet, and conquer the void.
          </p>
          <div className="flex flex-col gap-4 items-center w-full max-w-xs">
            <button 
              onClick={handleStart}
              disabled={isLoading}
              className="w-full px-8 py-4 bg-white text-black font-black italic uppercase tracking-tighter hover:bg-gray-200 transition-all transform hover:-translate-y-1 shadow-[0_0_20px_rgba(255,255,255,0.2)] disabled:opacity-50"
            >
              {user ? "Enter the Void" : (isLoading ? "Initializing..." : "Quick Start (Guest)")}
            </button>
            
            {!user && (
              <>
                <div className="flex items-center w-full gap-4">
                  <div className="h-[1px] bg-zinc-800 flex-1" />
                  <span className="text-[10px] text-zinc-600 font-bold uppercase tracking-widest">OR</span>
                  <div className="h-[1px] bg-zinc-800 flex-1" />
                </div>

                <button 
                  onClick={handleNeuralLink}
                  className="w-full px-8 py-4 bg-pink-600 hover:bg-pink-500 text-white font-bold uppercase tracking-widest transition-all transform hover:scale-105 shadow-[0_0_20px_rgba(236,72,153,0.4)]"
                >
                  Initialize Neural Link
                </button>
                
                <p className="text-[10px] text-zinc-500 uppercase tracking-[0.3em] text-center">
                  Social Login (Google/Email/Wallet)
                </p>
              </>
            )}
          </div>
        </div>

        {/* Decorative Elements */}
        <div className="absolute bottom-0 left-0 w-full h-64 bg-gradient-to-t from-pink-500/10 to-transparent pointer-events-none" />
      </main>
      
      <footer className="py-10 text-sm text-zinc-600 z-10">
        <p>© 2024 Project-0. All rights reserved.</p>
      </footer>
    </div>
  );
}
