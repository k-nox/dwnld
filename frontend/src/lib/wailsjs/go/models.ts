export namespace app {
	
	export enum Resolution {
	    RESOLUTION_2160 = "2160",
	    RESOLUTION_1440 = "1440",
	    RESOLUTION_1080 = "1080",
	    RESOLUTION_720 = "720",
	    RESOLUTION_480 = "480",
	    RESOLUTION_360 = "360",
	}
	export interface Options {
	    url: string;
	    outputDir?: string;
	    outputTempl?: string;
	    targetResolution?: Resolution;
	}

}

