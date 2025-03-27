import { Load } from '$lib/wailsjs/go/config/Manager';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ depends }) => {
	depends('app:defaults');
	return {
		defaults: await Load()
	};
};
