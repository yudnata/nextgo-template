'use server';

import { api } from '@/lib/api';
import { cookies } from 'next/headers';
import type { AuthUser } from './types';

type ActionState = {
  success: boolean;
  message?: string;
  error?: string;
  data?: unknown;
} | null;

export async function loginAction(state: ActionState, formData: FormData): Promise<ActionState> {
  const email = formData.get('email') as string;
  const password = formData.get('password') as string;

  try {
    const response = await api<{
      success: boolean;
      message?: string;
      data: { token: { access_token: string }; user: AuthUser };
    }>('/auth/login', {
      method: 'POST',
      body: JSON.stringify({ email, password }),
    });

    if (response.success) {
      const cookieStore = await cookies();
      cookieStore.set('auth_token', response.data.token.access_token, {
        httpOnly: true,
        secure: process.env.NODE_ENV === 'production',
        sameSite: 'lax',
        path: '/',
        maxAge: 60 * 60 * 24, // 24 hours
      });
      return { success: true };
    }

    return {
      success: false,
      error: response.message || 'Invalid credentials',
    };
  } catch (error) {
    return { success: false, error: (error as Error).message };
  }
}

export async function registerAction(state: ActionState, formData: FormData): Promise<ActionState> {
  const email = formData.get('email') as string;
  const password = formData.get('password') as string;
  const name = formData.get('name') as string;

  try {
    const response = await api<{ success: boolean }>('/auth/register', {
      method: 'POST',
      body: JSON.stringify({ email, password, name }),
    });

    return response;
  } catch (error) {
    return { success: false, error: (error as Error).message };
  }
}

export async function logoutAction() {
  const cookieStore = await cookies();
  cookieStore.delete('auth_token');
}
