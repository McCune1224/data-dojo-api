export const SITE = {
    title: "DataDojo Docs",
    description: "Official DataDojo API Documentation",
    defaultLanguage: "en-us",
} as const;

export const OPEN_GRAPH = {
    image: {
        src: "https://github.com/withastro/astro/blob/main/.github/assets/banner.png?raw=true",
        alt:
            "astro logo on a starry expanse of space," +
            " with a purple saturn-like planet floating in the right foreground",
    },
    twitter: "astrodotbuild",
};

export const KNOWN_LANGUAGES = {
    English: "en",
} as const;
export const KNOWN_LANGUAGE_CODES = Object.values(KNOWN_LANGUAGES);

export const GITHUB_EDIT_URL = `https://github.com/McCune1224/DataDojo/tree/main/docs`;

export const COMMUNITY_INVITE_URL = `https://discord.gg/hcXUjjjezw`;

// See "Algolia" section of the README for more information.
export const ALGOLIA = {
    indexName: "XXXXXXXXXX",
    appId: "XXXXXXXXXX",
    apiKey: "XXXXXXXXXX",
};

export type Sidebar = Record<
    (typeof KNOWN_LANGUAGE_CODES)[number],
    Record<string, { text: string; link: string }[]>
>;
export const SIDEBAR: Sidebar = {
    en: {
        Overview: [{ text: "Introduction", link: "en/introduction" }],
        Reference: [
            { text: "Overview", link: "en/overview" },
            { text: "Games", link: "en/games/overview" },
            { text: "Characters", link: "en/games/overview" },
            { text: "Moves", link: "en/games/overview" },
        ],
    },
};
