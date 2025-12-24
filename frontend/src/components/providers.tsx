'use client';

import {PrivyProvider} from '@privy-io/react-auth';
import {QueryClient, QueryClientProvider} from '@tanstack/react-query';
import {ReactNode, useState} from 'react';
import {useRouter} from 'next/navigation';

export default function Providers({children}: {children: ReactNode}) {
  const [queryClient] = useState(() => new QueryClient());
  const appId = process.env.NEXT_PUBLIC_PRIVY_APP_ID;
  const router = useRouter();

  if (!appId) {
    return (
      <div className="flex min-h-screen items-center justify-center bg-black text-white p-10 text-center">
        <div className="max-w-md space-y-4">
          <h1 className="text-2xl font-bold text-red-500">Configuration Missing</h1>
          <p className="text-zinc-400">
            <code>NEXT_PUBLIC_PRIVY_APP_ID</code> is not set in your environment variables.
          </p>
          <p className="text-sm text-zinc-500">
            Please create a project at <a href="https://dashboard.privy.io" className="text-indigo-400 underline">dashboard.privy.io</a> and add the App ID to your <code>.env.local</code> file.
          </p>
        </div>
      </div>
    );
  }

  return (
    <PrivyProvider
      appId={appId}
      onSuccess={() => {
        router.push('/hangar');
      }}
      config={{
        // Customize Privy's appearance and login methods
        appearance: {
          theme: 'dark',
          accentColor: '#676FFF',
          logo: 'https://your-logo-url.com',
        },
        // Create embedded wallets for users who don't have a wallet
        embeddedWallets: {
          ethereum: {
            createOnLogin: 'users-without-wallets',
          },
        },
        loginMethods: ['google', 'email', 'wallet', 'twitter'],
      }}
    >
      <QueryClientProvider client={queryClient}>
        {children}
      </QueryClientProvider>
    </PrivyProvider>
  );
}
