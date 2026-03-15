'use client';

import { useActionState } from 'react';
import { loginAction } from '../actions';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

export function LoginForm() {
  const [state, action, isPending] = useActionState(loginAction, null);

  return (
    <form
      action={action}
      className="space-y-4"
    >
      <div className="space-y-2">
        <label
          className="text-sm font-medium"
          htmlFor="email"
        >
          Email
        </label>
        <Input
          id="email"
          name="email"
          type="email"
          required
          placeholder="john@example.com"
        />
      </div>
      <div className="space-y-2">
        <label
          className="text-sm font-medium"
          htmlFor="password"
        >
          Password
        </label>
        <Input
          id="password"
          name="password"
          type="password"
          required
        />
      </div>

      {state?.error && <p className="text-sm text-red-500 font-medium">{state.error}</p>}

      <Button
        type="submit"
        className="w-full"
        disabled={isPending}
      >
        {isPending ? 'Logging in...' : 'Login'}
      </Button>
    </form>
  );
}
