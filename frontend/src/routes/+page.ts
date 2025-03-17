import { Defaults } from '$lib/wailsjs/go/download/Downloader';
import type { PageLoad } from './$types';

export const load: PageLoad = async () => {
	return {
		defaults: await Defaults()
	};
};
