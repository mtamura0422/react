


import React from 'react';
import { VariantProps, cva } from 'class-variance-authority';


export const ButtonVariants = cva("text-white rounded-md disabled:cursor-not-allowed disabled:bg-gray-300 disabled:hover:bg-gray-400", {
  variants: {
    intent: {
      primary: "bg-blue-400 hover:bg-blue-600",
      secondary: "bg-pink-400 hover:bg-pink-600",
    },
    size: {
      sm: "px-2 py-1 text-sm",
      md: "px-4 py-2 text-base",
      lg: "px-6 py-3 text-lg",
    },
  },
  defaultVariants: {
    intent: "primary",
    size: "md",
  },
});



export interface ButtonProps
  extends React.ButtonHTMLAttributes<HTMLButtonElement>,
    VariantProps<typeof ButtonVariants> {}


export const Button = React.forwardRef<HTMLButtonElement, ButtonProps>(
  (
    { className, intent, children, disabled, type = 'button', ...props },
    ref
  ) => {

  //  console.log("testing it");
    return (

      <button
        type={type}
        className={ButtonVariants({ intent, className })}
        ref={ref}
        disabled={disabled}
        {...props}
      >
        {children}
      </button>
    );
  }
);

export default Button;


