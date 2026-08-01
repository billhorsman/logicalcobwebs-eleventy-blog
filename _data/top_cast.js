import { loadTopFilms, topCast } from "../lib/filmsData.js";

export default () => topCast(loadTopFilms());
