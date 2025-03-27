import { type ClassValue, clsx } from 'clsx';
import { twMerge } from 'tailwind-merge';

export function cn(...inputs: ClassValue[]) {
	return twMerge(clsx(inputs));
}

export function inputClassNames<Type>(value: Type, defaultValue: Type) {
	return value && value !== defaultValue
		? 'border-indigo-400 focus:ring-indigo-400 focus-visible:ring-indigo-400'
		: '';
}
