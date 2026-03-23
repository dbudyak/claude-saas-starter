const API_BASE = import.meta.env.VITE_API_URL || '';

/**
 * Base API client. Attaches Authorization header if a token exists in localStorage.
 * Throws an Error with the server's message on non-2xx responses.
 */
export async function apiClient<T>(
  endpoint: string,
  options?: RequestInit
): Promise<T> {
  const token = localStorage.getItem('token');

  const response = await fetch(`${API_BASE}${endpoint}`, {
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
      ...options?.headers,
    },
  });

  if (!response.ok) {
    const text = await response.text().catch(() => response.statusText);
    const error = new Error(text || `Request failed: ${response.status}`) as Error & { status: number };
    error.status = response.status;
    throw error;
  }

  if (response.status === 204) {
    return undefined as T;
  }

  return response.json() as Promise<T>;
}
