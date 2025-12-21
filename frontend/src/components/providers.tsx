'use client';

import {PrivyProvider} from '@privy-io/react-auth';
import {QueryClient, QueryClientProvider} from '@tanstack/react-query';
import {ReactNode, useState} from 'react';

export default function Providers({children}: {children: ReactNode}) {
  const [queryClient] = useState(() => new QueryClient());

  return (
    <PrivyProvider
      appId={process.env.NEXT_PUBLIC_PRIVY_APP_ID || 'insert-your-privy-app-id-here'}
      config={{
        // Customize Privy's appearance and login methods
        appearance: {
          theme: 'dark',
          accentColor: '#676FFF',
          logo: 'https://your-logo-url.com',
        },
        // Create embedded wallets for users who don't have a wallet
        embeddedWallets: {
          createOnLogin: 'users-without-wallets',
        },
        loginMethods: ['email', 'wallet', 'google', 'twitter'],
      }}
    >
      <QueryClientProvider client={queryClient}>
        {children}
      </QueryClientProvider>
    </PrivyProvider>
  );
}
