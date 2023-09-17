import React, { ReactNode } from "react";

interface ButtonProps {
	/**
	 * unique identifier
	 */
	id: string;
	/**
	 * Button contents
	 */
	children: ReactNode;
	/**
	 * Button styling
	 */
	style?: keyof typeof getStyle;
	/**
	 * Button size
	 */
	size?: keyof typeof getSize;
	/**
	 * Additional classes
	 */
	className?: string;
	/**
	 * optional click handler
	 */
	onClick?: () => void;
}

const getStyle = {
	primary:
		"bg-kiokuDarkBlue border-kiokuDarkBlue text-eggshell shadow-sm hover:scale-105 hover:cursor-pointer",
	secondary:
		"bg-transparent border-transparent text-kiokuDarkBlue hover:scale-105 hover:cursor-pointer",
	error: "bg-kiokuRed border-kiokuRed text-white hover:scale-105 hover:cursor-pointer",
	warning:
		"bg-kiokuYellow border-kiokuYellow text-white hover:scale-105 hover:cursor-pointer",
	disabled:
		"bg-gray-500 border-gray-500 text-eggshell hover:cursor-not-allowed",
} as const;

const getSize = {
	sm: "px-3 py-1.5 text-xs sm:text-xs md:text-sm lg:px-3 lg:py-1.5 lg:text-base xl:text-lg",
	md: "px-3 py-1.5 text-xs sm:text-sm md:text-base lg:px-5 lg:py-3 lg:text-lg xl:text-xl",
	lg: "px-5 py-3 text-sm sm:text-base md:text-lg lg:px-5 lg:py-3 lg:text-xl xl:text-2xl",
} as const;

/**
 * UI component for user interactions
 */
export const Button = ({
	className,
	style = "primary",
	size = "md",
	children = "",
	...props
}: ButtonProps) => {
	return (
		<button
			className={`flex items-center justify-center rounded-md border-2 text-center font-semibold outline-none transition ${getStyle[style]} ${getSize[size]} ${className}`}
			{...props}
		>
			{children}
		</button>
	);
};
