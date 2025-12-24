"use client";

import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { useAuthSync } from "@/hooks/use-auth-sync";

export default function CharacterCreation() {
  const router = useRouter();
  const { user: backendUser, isLoading: isSyncing } = useAuthSync();
  const [name, setName] = useState("");
  const [gender, setGender] = useState("MALE");
  const [faceIndex, setFaceIndex] = useState(0);
  const [hairIndex, setHairIndex] = useState(0);
  const [isSubmitting, setIsSubmitting] = useState(false);

  useEffect(() => {
    if (!isSyncing && !backendUser) {
      router.push('/');
    } else if (!isSyncing && backendUser?.active_character_id) {
      router.push('/hangar');
    }
  }, [isSyncing, backendUser, router]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!name) return;

    setIsSubmitting(true);
    try {
      const token = localStorage.getItem("project0_token");
      const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/v1/auth/characters/create`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`,
        },
        body: JSON.stringify({
          name,
          gender,
          face_index: faceIndex,
          hair_index: hairIndex,
        }),
      });

      if (response.ok) {
        // Refresh user data or just redirect
        window.location.href = "/hangar";
      } else {
        const err = await response.text();
        alert(`Failed to create character: ${err}`);
      }
    } catch (error) {
      console.error("Error creating character:", error);
    } finally {
      setIsSubmitting(false);
    }
  };

  if (isSyncing || !backendUser) {
    return <div className="flex min-h-screen items-center justify-center bg-black text-white">Loading...</div>;
  }

  return (
    <div className="flex min-h-screen flex-col items-center justify-center p-6">
      <div className="w-full max-w-md space-y-8 bg-black/60 backdrop-blur-xl p-8 rounded-xl border border-zinc-800 shadow-2xl">
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
            <label className="block text-sm font-medium text-zinc-400 uppercase tracking-wider mb-2">Gender</label>
            <div className="flex gap-4">
              {["MALE", "FEMALE", "NON_BINARY"].map((g) => (
                <button
                  key={g}
                  type="button"
                  onClick={() => setGender(g)}
                  className={`flex-1 py-2 rounded-md border transition-all ${
                    gender === g ? "bg-pink-500 border-pink-500 text-white" : "bg-black border-zinc-700 text-zinc-400 hover:border-zinc-500"
                  }`}
                >
                  {g.replace("_", " ")}
                </button>
              ))}
            </div>
          </div>

          <div className="grid grid-cols-2 gap-4">
            <div>
              <label className="block text-sm font-medium text-zinc-400 uppercase tracking-wider mb-2">Face Type</label>
              <div className="flex items-center gap-2">
                <button
                  type="button"
                  onClick={() => setFaceIndex(Math.max(0, faceIndex - 1))}
                  className="p-2 bg-zinc-800 rounded-md hover:bg-zinc-700"
                >
                  &lt;
                </button>
                <span className="flex-1 text-center font-mono">{faceIndex + 1}</span>
                <button
                  type="button"
                  onClick={() => setFaceIndex(faceIndex + 1)}
                  className="p-2 bg-zinc-800 rounded-md hover:bg-zinc-700"
                >
                  &gt;
                </button>
              </div>
            </div>
            <div>
              <label className="block text-sm font-medium text-zinc-400 uppercase tracking-wider mb-2">Hair Style</label>
              <div className="flex items-center gap-2">
                <button
                  type="button"
                  onClick={() => setHairIndex(Math.max(0, hairIndex - 1))}
                  className="p-2 bg-zinc-800 rounded-md hover:bg-zinc-700"
                >
                  &lt;
                </button>
                <span className="flex-1 text-center font-mono">{hairIndex + 1}</span>
                <button
                  type="button"
                  onClick={() => setHairIndex(hairIndex + 1)}
                  className="p-2 bg-zinc-800 rounded-md hover:bg-zinc-700"
                >
                  &gt;
                </button>
              </div>
            </div>
          </div>

          <button
            type="submit"
            disabled={isSubmitting || !name}
            className="w-full bg-pink-600 hover:bg-pink-500 text-white font-bold py-3 rounded-md uppercase tracking-widest transition-all disabled:opacity-50 disabled:cursor-not-allowed mt-4 shadow-lg shadow-pink-500/20"
          >
            {isSubmitting ? "Registering..." : "Initialize Pilot"}
          </button>
        </form>
      </div>
    </div>
  );
}
