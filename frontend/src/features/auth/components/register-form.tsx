'use client';

import { useActionState } from 'react';
import { registerAction } from '../actions';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

const initialState = {
  success: false,
  message: '',
  error: '',
};

export function RegisterForm() {
  const [state, action, isPending] = useActionState(registerAction, initialState);

  return (
    <form
      action={action}
      className="space-y-4"
    >
      <div className="space-y-2">
        <label
          className="text-sm font-medium"
          htmlFor="name"
        >
          Name
        </label>
        <Input
          id="name"
          name="name"
          required
          placeholder="John Doe"
        />
      </div>
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
          minLength={8}
        />
      </div>

      {state?.error && <p className="text-sm text-red-500 font-medium">{state.error}</p>}

      {state?.success && (
        <p className="text-sm text-green-500 font-medium">
          Registration successful! You can now log in.
        </p>
      )}

      <Button
        type="submit"
        className="w-full"
        disabled={isPending}
      >
        {isPending ? 'Creating account...' : 'Register'}
      </Button>
    </form>
  );
}
