import { describe, it, expect, vi, beforeEach } from 'vitest';
import { apiClient } from './client';

describe('apiClient', () => {
  beforeEach(() => {
    vi.stubGlobal('fetch', vi.fn());
    localStorage.clear();
  });

  it('includes Authorization header when token exists', async () => {
    localStorage.setItem('token', 'test-token');

    const mockResponse = { ok: true, status: 200, json: async () => ({ data: 'ok' }) };
    vi.mocked(fetch).mockResolvedValueOnce(mockResponse as Response);

    await apiClient('/api/test');

    const call = vi.mocked(fetch).mock.calls[0];
    const headers = call[1]?.headers as Record<string, string>;
    expect(headers['Authorization']).toBe('Bearer test-token');
  });

  it('throws on non-ok response', async () => {
    const mockResponse = { ok: false, status: 400, text: async () => 'bad request' };
    vi.mocked(fetch).mockResolvedValueOnce(mockResponse as Response);

    await expect(apiClient('/api/test')).rejects.toThrow('bad request');
  });
});
