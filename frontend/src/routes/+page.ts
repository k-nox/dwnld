import { Settings } from '$lib/wailsjs/go/app/App';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ depends }) => {
	depends('app:defaults');
	return {
		defaults: await Settings()
	};
};
