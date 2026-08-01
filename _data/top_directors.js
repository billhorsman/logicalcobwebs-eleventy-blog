import { loadTopFilms, topDirectors } from "../lib/filmsData.js";

export default () => topDirectors(loadTopFilms());
