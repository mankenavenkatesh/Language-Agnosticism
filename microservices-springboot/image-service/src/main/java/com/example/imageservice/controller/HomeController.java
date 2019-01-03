package com.example.imageservice.controller;

import java.util.Arrays;
import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.env.Environment;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.example.imageservice.entities.Image;



@RestController
@RequestMapping("/")
public class HomeController {
	@Autowired
	private Environment env;
		
	@RequestMapping("/images")
	public List<Image> getImages() throws Exception {
		List<Image> images = Arrays.asList(
			new Image(1, "Treehouse of Horror V", "https://www.imdb.com/title/tt0096697/mediaviewer/rm3842005760"),
			new Image(2, "The Town", "https://www.imdb.com/title/tt0096697/mediaviewer/rm3698134272"),
			new Image(3, "The Last Traction Hero", "https://www.imdb.com/title/tt0096697/mediaviewer/rm1445594112"));
		
		throw new Exception("Images can't be fetched");
//		return images;
	}
}
