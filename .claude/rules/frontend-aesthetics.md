# Frontend Aesthetics

Avoid generic "AI slop" aesthetics. Every product deserves a distinctive visual identity.

## Typography

- Choose fonts that are beautiful and distinctive
- **Avoid**: Inter, Roboto, Arial, system-ui as primary fonts
- **Better**: DM Sans, Outfit, Plus Jakarta Sans, Syne, Cabinet Grotesk, General Sans
- Import from Bunny Fonts (GDPR-friendly) or Google Fonts
- Vary between projects — don't default to the same font every time

## Color

- Commit to a cohesive palette — don't pick random colors
- Use CSS custom properties for consistency:
  ```css
  :root {
    --color-primary: #1a1a2e;
    --color-accent: #e94560;
    --color-surface: #16213e;
  }
  ```
- Dominant color + sharp accent outperforms timid, even distribution
- **Avoid**: Purple gradients on white (most overused AI aesthetic)
- Draw from context: B2B feels different from consumer, fintech from gaming

## Motion

- CSS transitions on interactive elements: `transition-colors duration-200`
- Page load: one staggered reveal beats scattered micro-interactions
- Use `animation-delay` for staggered entry:
  ```css
  .card:nth-child(1) { animation-delay: 0ms; }
  .card:nth-child(2) { animation-delay: 100ms; }
  ```
- Framer Motion for complex animations in React

## Backgrounds

- Layer CSS gradients — don't default to solid colors
- Subtle grid or dot patterns add depth
- Dark themes with colored accents often feel more premium

## Layout

- Create visual hierarchy — not everything at the same weight
- Use whitespace intentionally
- Cards with subtle borders beat heavy shadows in 2024+

## Anti-Patterns

- ❌ Generic blue/purple gradient hero sections
- ❌ Overused card-with-shadow pattern everywhere
- ❌ Inter font on everything
- ❌ "Dashboard purple" color scheme
- ❌ Same layout as every other SaaS landing page

## Principle

Make unexpected choices that feel genuinely designed for the specific context.
Every project should have its own visual personality.
