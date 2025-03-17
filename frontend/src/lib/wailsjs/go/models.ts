export namespace config {
	
	export enum Resolution {
	    RESOLUTION_2160 = "2160",
	    RESOLUTION_1440 = "1440",
	    RESOLUTION_1080 = "1080",
	    RESOLUTION_720 = "720",
	    RESOLUTION_480 = "480",
	    RESOLUTION_360 = "360",
	}
	export interface Download {
	    outputDirectory?: string;
	    outputTemplate?: string;
	    targetResolution?: Resolution;
	}
	export interface Settings {
	    download: Download;
	}

}

