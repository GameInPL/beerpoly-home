import '../styles/index.scss';
import * as FragmentComponent from './components/fragment/component';
import Parallax from 'parallax-js'

document.addEventListener("DOMContentLoaded", () => {
  let body = document.getElementsByTagName("BODY")[0];
  FragmentComponent.init(body);
});

document.addEventListener("DOMContentLoaded", () => {
  let scene = document.getElementById('scene');
  let parallax = new Parallax(scene);
});
