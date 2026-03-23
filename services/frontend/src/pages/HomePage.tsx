import { Link } from 'react-router-dom';

export default function HomePage() {
  return (
    <div className="min-h-screen flex flex-col items-center justify-center px-4"
         style={{ background: 'var(--color-bg)' }}>
      <div className="text-center max-w-lg">
        <h1 className="text-5xl font-bold mb-4" style={{ color: 'var(--color-text)' }}>
          SaaS Starter
        </h1>
        <p className="text-lg mb-8" style={{ color: 'var(--color-text-muted)' }}>
          Go + React. Production-ready. Claude Code optimized.
        </p>
        <div className="flex gap-4 justify-center">
          <Link
            to="/signup"
            className="px-6 py-3 rounded-lg font-medium transition-colors"
            style={{ background: 'var(--color-primary)', color: '#fff' }}
          >
            Get Started
          </Link>
          <Link
            to="/login"
            className="px-6 py-3 rounded-lg font-medium transition-colors"
            style={{
              background: 'var(--color-surface)',
              color: 'var(--color-text)',
              border: '1px solid var(--color-border)',
            }}
          >
            Sign In
          </Link>
        </div>
      </div>
    </div>
  );
}
