const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

type RequestOptions = RequestInit & {
  params?: Record<string, string>;
};

/**
 * API client for communicating with the Go backend.
 * Wraps fetch with base URL, JSON headers, and error handling.
 */
export async function api<T>(
  endpoint: string,
  options: RequestOptions = {}
): Promise<T> {
  const { params, ...fetchOptions } = options;

  const url = new URL(`${API_BASE_URL}${endpoint}`);
  if (params) {
    Object.entries(params).forEach(([key, value]) =>
      url.searchParams.append(key, value)
    );
  }

  const headers = new Headers(options.headers);
  headers.set("Content-Type", "application/json");

  // Add token if exists
  let token: string | null = null;
  
  if (typeof window !== "undefined") {
    // Client-side: use localStorage
    token = localStorage.getItem("auth_token");
  } else {
    // Server-side: use cookies (Note: this only works in Server Components/Actions)
    try {
      const { cookies } = await import("next/headers");
      const cookieStore = await cookies();
      token = cookieStore.get("auth_token")?.value || null;
    } catch {
      // Not in a server context that supports cookies
    }
  }

  if (token) {
    headers.set("Authorization", `Bearer ${token}`);
  }

  const response = await fetch(url.toString(), {
    ...fetchOptions,
    headers,
  });

  if (!response.ok) {
    throw new Error(`API Error: ${response.status} ${response.statusText}`);
  }

  return response.json() as Promise<T>;
}