'use client';

import { useState } from "react";
import { useAuthSync } from "@/hooks/use-auth-sync";
import { pilotSystem } from "@/systems/PilotSystem";

interface CharacterCreationStageProps {
  onSuccess: () => void;
}

export default function CharacterCreationStage({ onSuccess }: CharacterCreationStageProps) {
  const { refreshUser } = useAuthSync();
  const [name, setName] = useState("");
  const [gender, setGender] = useState("MALE");
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!name) return;

    setIsSubmitting(true);
    try {
      await pilotSystem.createCharacter({
        name,
        gender,
        face_index: 0,
        hair_index: 0,
      });
      
      await refreshUser();
      onSuccess();
    } catch (error: any) {
      console.error("Error creating character:", error);
      alert(`Failed to create character: ${error.message}`);
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="flex min-h-screen flex-col items-center justify-center p-6 bg-black text-white">
      <div className="w-full max-w-md space-y-8 bg-zinc-900/40 backdrop-blur-xl p-8 rounded-xl border border-zinc-800 shadow-2xl">
        <div className="text-center">
          <h1 className="text-3xl font-bold tracking-tighter uppercase italic text-pink-500">Pilot Registration</h1>
          <p className="text-zinc-400 mt-2">Create your identity in the void.</p>
        </div>

        <form onSubmit={handleSubmit} className="space-y-6">
          <div>
            <label className="block text-sm font-medium text-zinc-400 uppercase tracking-wider mb-2">Pilot Name</label>
            <input
              type="text"
              value={name}
              onChange={(e) => setName(e.target.value)}
              className="w-full bg-black border border-zinc-700 rounded-md px-4 py-2 focus:outline-none focus:border-pink-500 transition-colors"
              placeholder="Enter name..."
              required
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-zinc-400 uppercase tracking-wider mb-4 text-center">Select Your Pilot Identity</label>
            <div className="flex gap-6">
              {[
                { id: "MALE", label: "Male" },
                { id: "FEMALE", label: "Female" }
              ].map((g) => (
                <button
                  key={g.id}
                  type="button"
                  onClick={() => setGender(g.id)}
                  className={`flex-1 group relative rounded-xl overflow-hidden border-2 p-4 transition-all duration-300 ${
                    gender === g.id 
                      ? "border-pink-500 shadow-[0_0_20px_rgba(236,72,153,0.3)] scale-105" 
                      : "border-zinc-800 hover:border-zinc-600"
                  }`}
                >
                  <div className="text-center font-bold uppercase tracking-widest text-xs">{g.label}</div>
                </button>
              ))}
            </div>
          </div>

          <button
            type="submit"
            disabled={isSubmitting || !name}
            className="w-full py-4 bg-pink-600 hover:bg-pink-500 disabled:bg-zinc-800 disabled:text-zinc-500 text-white font-bold uppercase tracking-widest transition-all rounded-md"
          >
            {isSubmitting ? "Registering..." : "Complete Registration"}
          </button>
        </form>
      </div>
    </div>
  );
}
