import { render, screen } from '@hitz-group/monorepo-react-testing/react';
import React from 'react';
import App from './App';

describe('App', () => {
  test('should render counter with 0', () => {
    render(<App />);
    expect(screen.getByRole('button')).toHaveTextContent('count is: 0');
  });
});
