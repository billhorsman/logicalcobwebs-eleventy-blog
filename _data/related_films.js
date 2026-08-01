import { loadTopFilms, relatedFilms } from "../lib/filmsData.js";

export default () => relatedFilms(loadTopFilms());
