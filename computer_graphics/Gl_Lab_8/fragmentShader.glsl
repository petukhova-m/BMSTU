#version 450 core

		out vec4 FragColor;
		
		in vec2 TexCoords;
		in vec3 NormalCoords;
		in vec3 gPos;
		uniform sampler2D texture1;
		uniform sampler2D spec_map;
		uniform sampler2D amb_map;
		uniform int isText;
		uniform vec3 cylCol;
		uniform vec3 lightCol;
		uniform vec3 lightPos;
		uniform vec3 projDir;
		uniform vec3 viewPos;
		uniform mat4 projRot;
		uniform float exp;
		uniform float angle;
		
		void main()
		{
			vec3 tex = vec3(texture(texture1, TexCoords));
			vec3 sMap = vec3(texture(spec_map, TexCoords));
			vec3 aMap = vec3(texture(amb_map, TexCoords));

			//Diffuse
			vec3 n = normalize(NormalCoords);
			vec3 proj = normalize(projDir);
			vec3 lightDir = normalize(lightPos - gPos);
			lightDir = vec3(projRot * vec4(lightDir, 1.0));
			float diffEff = max(dot(n, lightDir), 0.0);
			float projAngle = max(dot(-lightDir, proj), 0.0);
			if (degrees(acos(diffEff)) > angle) diffEff = 0;
			vec3 diffuse = pow(projAngle, exp) * diffEff * lightCol * 0.8;
			if (isText == 0) diffuse = tex * diffuse;

            //Ambient
			float ambEff = 0.25;
			vec3 ambient = ambEff * lightCol;
			if (isText == 0) ambient = aMap * ambient;

			//Specular
			float specEff = 0.9;
			vec3 viewDir = normalize(viewPos - gPos);
			vec3 reflDir = reflect(-lightDir, n);
			float spec = pow(max(dot(viewDir, reflDir), 0.0), exp/4);
			vec3 specular;
			if (isText == 0) specular = sMap * specEff * spec * lightCol;
			else specular = vec3(0);

			vec3 res = (ambient + diffuse + specular) * cylCol;
			FragColor = vec4(res, 1.0);
		}
