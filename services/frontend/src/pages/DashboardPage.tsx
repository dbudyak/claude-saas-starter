import { useQuery } from '@tanstack/react-query';
import { useNavigate } from 'react-router-dom';
import { getMe } from '../api/auth';
import { useAuthStore } from '../store/authStore';

export default function DashboardPage() {
  const navigate = useNavigate();
  const logout = useAuthStore((s) => s.logout);

  const { data: user, isLoading } = useQuery({
    queryKey: ['me'],
    queryFn: getMe,
  });

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  return (
    <div className="min-h-screen px-6 py-8" style={{ background: 'var(--color-bg)' }}>
      <div className="max-w-2xl mx-auto">
        <div className="flex items-center justify-between mb-8">
          <h1 className="text-2xl font-bold" style={{ color: 'var(--color-text)' }}>
            Dashboard
          </h1>
          <button
            onClick={handleLogout}
            className="px-4 py-2 rounded-lg text-sm font-medium transition-colors"
            style={{
              background: 'var(--color-surface)',
              border: '1px solid var(--color-border)',
              color: 'var(--color-text-muted)',
              cursor: 'pointer',
            }}
          >
            Sign Out
          </button>
        </div>

        {isLoading ? (
          <p style={{ color: 'var(--color-text-muted)' }}>Loading...</p>
        ) : user ? (
          <div className="p-6 rounded-xl" style={{
            background: 'var(--color-surface)',
            border: '1px solid var(--color-border)',
          }}>
            <p className="text-sm mb-1" style={{ color: 'var(--color-text-muted)' }}>Signed in as</p>
            <p className="font-medium" style={{ color: 'var(--color-text)' }}>{user.email}</p>
            <p className="text-sm mt-1" style={{ color: 'var(--color-text-muted)' }}>
              Role: {user.role}
            </p>
          </div>
        ) : null}

        <div className="mt-8 p-6 rounded-xl" style={{
          background: 'var(--color-surface)',
          border: '1px solid var(--color-border)',
        }}>
          <h2 className="font-semibold mb-2" style={{ color: 'var(--color-text)' }}>
            Your app starts here
          </h2>
          <p className="text-sm" style={{ color: 'var(--color-text-muted)' }}>
            Add your features to this page. The auth flow is wired up — JWT token is stored in localStorage
            and sent automatically with every API request.
          </p>
        </div>
      </div>
    </div>
  );
}
