# React & TypeScript Standards

## Naming

- **Components**: PascalCase (`UserProfile.tsx`)
- **Hooks**: camelCase with `use` prefix (`useAuth.ts`)
- **Utilities**: camelCase (`dateUtils.ts`)
- **Types**: PascalCase (`User`, `ApiResponse`)
- **Constants**: SCREAMING_SNAKE_CASE (`API_BASE_URL`)

## Component Pattern

```tsx
import { FC } from 'react';

interface Props {
  userId: string;
  onUpdate?: (id: string) => void;
}

export const UserCard: FC<Props> = ({ userId, onUpdate }) => {
  const { data: user, isLoading } = useQuery({
    queryKey: ['user', userId],
    queryFn: () => getUser(userId),
  });

  if (isLoading) return <Spinner />;
  if (!user) return null;

  return (
    <div className="p-4 rounded-lg border">
      <h3>{user.name}</h3>
    </div>
  );
};
```

## TypeScript Rules

- Never use `any` — use `unknown` or specific types
- Always type component props with interface
- Use union types for enums: `'admin' | 'member'`
- Type useState explicitly: `useState<User | null>(null)`

## State Management

- **Server state**: React Query (`useQuery`, `useMutation`)
- **UI state**: Zustand stores
- **Form state**: local `useState` (or React Hook Form for complex forms)
- **Never** use Redux — overkill for SaaS apps of this scale

## React Query Patterns

```tsx
// Fetching
const { data, isLoading, error } = useQuery({
  queryKey: ['items', filter],
  queryFn: () => getItems(filter),
});

// Mutating with cache invalidation
const { mutate } = useMutation({
  mutationFn: createItem,
  onSuccess: () => {
    queryClient.invalidateQueries({ queryKey: ['items'] });
  },
});
```

## API Client

```typescript
// Never put API calls directly in components
// src/api/items.ts
export const getItems = (): Promise<Item[]> =>
  apiClient('/api/items');

export const createItem = (data: CreateItemRequest): Promise<Item> =>
  apiClient('/api/items', { method: 'POST', body: JSON.stringify(data) });
```

## Styling (TailwindCSS)

- Use utility classes directly — no CSS files for component styles
- Extract repeated patterns to components, not CSS classes
- Use `clsx` or `cn` for conditional classes:
  ```tsx
  className={clsx('base-classes', condition && 'conditional-class')}
  ```

## Performance

```tsx
// Memoize expensive components
export const ItemCard = memo<Props>(({ item }) => <div>{item.name}</div>);

// Memoize expensive computations
const sorted = useMemo(() => items.sort(...), [items]);

// Stable callbacks
const handleClick = useCallback(() => navigate(`/items/${id}`), [id]);
```

## Error Handling

```tsx
// Always handle loading and error states
if (isLoading) return <LoadingState />;
if (error) return <ErrorState message="Failed to load" />;
if (!data?.length) return <EmptyState />;
```

## Anti-Patterns

- ❌ Using `any` type
- ❌ Mutating state directly
- ❌ Missing useEffect dependencies
- ❌ Not cleaning up subscriptions
- ❌ Prop drilling beyond 2 levels (use context or Zustand)
- ❌ Ignoring TypeScript errors with `@ts-ignore`
- ❌ API calls directly in components (use api/ functions)
