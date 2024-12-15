# Sliderz

Sliderz is a ChatGPT wrapper, where you can upload a file and it generates the slides and notes you need for your presentation.

## Why/What/How?

-   **Why?** Because we all have been there, where we have to present a report and we have to make slides for it. It's a tedious task to make slides and notes for the slides. So, why not automate it?

-   **What?** Sliderz is a ChatGPT wrapper, where you can upload a file and it generates the slides and notes you need for your presentation.

-   **How?** You can upload a file and it will generate the slides and notes for you. You can also customize the number of slides you want and the length of the notes.

### Why make it tho?

This is a learning project for me. I am trying to learn new technologies to see if they are useful. The main focuses of the project is:

-   Learning go, for the backend
-   ~~Learning SolidJS, for the frontend~~ I tried, but solid seems confusing with typescript!
-   Learning Svelte, for the frontend

Thas pretty much it :)

## How to run?

Run the backend:

```bash
cd backend && go run .
// Or live reload
cd backend && air
```

Run the frontend:

```bash
cd frontend && npm i && npm run dev
```

## Where are we now?

-   [x] Frontend setup
-   [x] Backend setup
-   [ ] Simple frontend
-   [ ] Simple backend
-   [ ] File upload
-   [ ] Parse the file
-   [ ] Generate slides from parsed sections, using ChatGPT
-   [ ] Generate notes from parsed sections, using ChatGPT
-   [ ] Return the slides and notes
-   [ ] Make a nice frontend
