/// <reference types="@sveltejs/kit" />

export {};

interface UserInterface {
  id: string;
  username: string;
  disc: string;
  avatar: string;
}

interface PageData {
  user: UserInterface;
  token: string;
  css: string?;
  refresh_token: string?;
  user_experiments: number[];
  site_lang: string;
}

interface UserFnInterface {
  id: string;
  token: string;
  apiUrl: string;
  lynxUrl: string;
}

declare global {
  namespace App {
    interface Locals {}

    interface Platform {}

    interface PageData { 
      user: UserInterface;
      token: string;
      css: string?;
      refresh_token: string?;
      user_experiments: number[];
      site_lang: string;
      admin_ticket: string;
      url: string
    };

    interface Stuff {}
  }

  namespace globalThis {
    interface Window {
      user: () => UserFnInterface;
      llhandler: () => any;
      QuillMarkdown: any;
    }
  }
}
