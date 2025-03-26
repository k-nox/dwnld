export interface FieldProps<Type> {
	value?: Type;
	defaultValue?: Type;
	disabled: boolean;
}

export interface LabeledFieldProps<Type> extends FieldProps<Type> {
	label: string;
}
