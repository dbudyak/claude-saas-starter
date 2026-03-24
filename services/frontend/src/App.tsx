import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import HomePage from './pages/HomePage';

// Add your routes here.
// For protected routes, wrap with an auth guard component that reads from your auth store.
const router = createBrowserRouter([
  {
    path: '/',
    element: <HomePage />,
  },
]);

export default function App() {
  return <RouterProvider router={router} />;
}
