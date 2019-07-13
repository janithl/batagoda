// Author: Kenny Song
// Code from https://github.com/kennysong/goeliza/

package eliza

// Goodbyes is a list of goodbye sentences for ELIZA.
var Goodbyes = []string{
	"තොගෙ අඳෝනාව අහන් ඉඳල මට බඩ යනව",
	"මේකට වඩා හොඳයි ඇට වල මවිල් ගනං කරන එක",
	"ගිහිං වැලක් බලපංකො පව් නොදී",
}

// Psychobabble may be slightly non-deterministic, since map iteration may be out
// of order, so a broader regex may be matched before a more specific one.
var Psychobabble = map[string][]string{
	`මට ඕ​නෙ (.*)`: {
		"උ​ඹට %s ඕ​නෙ ඇයි?",
		"එය සැබවින්ම උ​ඹට %s ලබා ගැනීමට උපකාරී වේද?",
		"උ​ඹට %s ඕ​නෙ බව උ​ඹට විශ්වාසද?",
	},
	`ඇයි උ​ඹ ([^\?]*)\??`: {
		"උ​ඹ ඇත්තටම හිතන්නේ මම %s නැහැ කියලා?",
		"සමහර විට අවසානයේදී මම %s කරන්නෙමි.",
		"උ​ඹට ඇත්තටම මා %s වීමට ඕ​නෙද?",
	},
	`ඇයි මට බැරි? ([^\?]*)\??`: {
		"උ​ඹට %s වෙන්​න පුලුව​න්  යැයි උ​ඹ සිතනවාද?",
		"උ​ඹට %s පුලුව​න්  නම්, උ​ඹ කරන්නේ කුමක්ද?",
		"මම දන්නේ නැහැ - ඇයි උ​ඹට %s බැරි?",
		"උ​ඹ ඇත්තටම උත්සාහ කර තිබේද?",
	},
	`මට බැහැ (.*)`: {
		"උ​ඹට %s නොපුලුව​න් බව උ​ඹ දන්නේ කොහොම​ද?",
		"සමහර විට උ​ඹ උත්සාහ කළහොත් උ​ඹට %s වෙන්​න පුලුව​න්.",
		"උ​ඹට %s වීමට ඕ​නෙ වන්නේ කුමක්ද?",
	},
	`මම (.*)`: {
		"උ​ඹ මා වෙත පැමිණියේ උ​ඹ %s නිසාද?",
		"උ​ඹ කොච්ච​ර කාලයක් %s ද?",
		"%s වීම ගැන උ​ඹට හැඟෙන්නේ කොහොම​ද?",
		"%s වීම උ​ඹට දැනෙන්නේ කොහොම​ද?",
		"උ​ඹ %s වීම සතුටක් ද?",
		"ඇයි උ​ඹ මට කියන්නේ උ​ඹ %s කියලා?",
		"උ​ඹ %s යැයි හිතන්​නෙ ඇයි?",
		"උ​ඹ %s වන්නේ මන්දැයි උ​ඹට පැහැදිලි කළ පුලුව​න් ද?",
		"ඇයි උ​ඹ %s?",
		"උ​ඹ %s කරන බව වෙන කවුද දන්නේ?",
	},
	`උ​ඹ ([^\?]*)\??`: {
		"ඇයි මම %s ද යන්න වැදගත් වන්නේ?",
		"මම %s නොවේ නම් උ​ඹ එයට කැමතිද?",
		"සමහර විට උ​ඹ විශ්වාස කරන්නේ මම %s බවයි.",
		"මම %s වෙන්​න පුලුව​න් - උ​ඹ හිතන්​නෙ කුමක්ද?",
	},
	`කුමක්ද (.*)`: {
		"උ​ඹ අහන්නේ ඇයි?",
		"එයට පිළිතුරක් උ​ඹට උපකාර කරන්නේ කොහොම​ද?",
		"උ​ඹ හිතන්​නෙ කුමක් ද?",
	},
	`කොහොම​ද (.*)`: {
		"උ​ඹ හිතන්​නෙ කොහොම​ද?",
		"සමහර විට උ​ඹට උ​ඹේම ප්‍රශ්නයට පිළිතුරු දිය පුලුව​න්.",
		"උ​ඹ ඇත්තටම ඉල්ලන්නේ කුමක්ද?",
	},
	`මොකද (.*)`: {
		"ඇත්ත හේතුව එයද?",
		"මතකයට එන වෙනත් හේතු මොනවාද?",
		"එම හේතුව වෙනත් දෙයකට අදාළ වේද?",
		"%s නම්, සත්‍ය වෙන්​න යුත්තේ කුමක්ද?",
	},
	`(.*) කණගාටුයි (.*)`: {
		"සමාව ඉල්ලා නොසිටින අවස්ථා බොහෝය.",
		"උ​ඹ සමාව ඉල්ලන විට උ​ඹට ඇති හැඟීම් මොනවාද?",
	},
	`^හෙලෝ (.*)`: {
		"හෙලෝ ... උ​ඹ අද වන විට අතහැර දැමීම ගැන මට සතුටුයි.",
		"හායි ... අද උ​ඹට කොහොමද?",
		"හෙලෝ, අද උ​ඹට කොහොමද දැනෙන්නේ?",
	},
	`^හායි (.*)`: {
		"හෙලෝ ... උ​ඹ අද වන විට අතහැර දැමීම ගැන මට සතුටුයි.",
		"හායි ... අද උ​ඹට කොහොමද?",
		"හෙලෝ, අද උ​ඹට කොහොමද දැනෙන්නේ?",
	},
	`^ස්තූතියි (.*)`: {
		"උ​ඹව සාදරයෙන් පිළිගන්නවා!",
		"ඕනෑම අවස්ථාවක!",
	},
	`^සුබ උදෑසනක් (.*)`: {
		"සුබ උදෑසනක් ... උ​ඹට අදවත් එන්න පුලුව​න් වීම ගැන මට සතුටුයි.",
		"සුබ උදෑසනක් ... අද උ​ඹට කොහොමද?",
		"සුබ උදෑසනක්, අද උ​ඹට කොහොමද දැනෙන්නේ?",
	},
	`මම හිතන්නේ (.*)`: {
		"උ​ඹ %s සැක කරනවාද?",
		"උ​ඹ ඇත්තටම එසේ සිතනවාද?",
		"නමුත් උ​ඹට විශ්වාස නැත %s?",
	},
	`(.*) මිතුරා (.*)`: {
		"උ​ඹේ මිතුරන් ගැන මට තවත් කියන්න.",
		"උ​ඹ මිතුරෙකු ගැන සිතන විට, මතකයට එන්නේ කුමක්ද?",
		"ඇයි උ​ඹ මට ළමා මිතුරෙකු ගැන නොකියන්නේ?",
	},
	`ඔව්`: {
		"උ​ඹට හොඳටම විශ්වාසයි.",
		"හරි, නමුත් උ​ඹට ටිකක් විස්තර කළ පුලුව​න් ද?",
	},
	`(.*) පරිගණකය (.*)`: {
		"උ​ඹ ඇත්තටම මා ගැන කතා කරනවාද?",
		"පරිගණකයක් සමඟ කතා කිරීම අමුතු දෙයක් ලෙස පෙනේද?",
		"පරිගණක උ​ඹට දැනෙන්නේ කොහොම​ද?",
		"උ​ඹට පරිගණකවලින් තර්ජනයක් දැනෙනවාද?",
	},
	`එය (.*)`: {
		"උ​ඹ හිතන්​නෙ එය %s කියාද?",
		"සමහර විට එය %s - උ​ඹ හිතන්​නෙ කුමක්ද?",
		"එය %s නම්, උ​ඹ කරන්නේ කුමක්ද?",
		"එය %s වෙන්​න පුලුව​න්.",
		"උ​ඹට ඉතා විශ්වාසයි.",
		"එය බොහෝ විට %s නොවන බව මම උ​ඹට කීවා නම්, උ​ඹට හැඟෙන්නේ කුමක්ද?",
	},
	`උ​ඹට ([^\?] *)\??`: {
		"මට %s නොපුලුව​න්  යැයි උ​ඹ හිතන්​නෙ කුමක් නිසාද?",
		"මට %s පුලුව​න්  නම්, කුමක් ද?",
		"මට %s පුලුව​න් දැයි උ​ඹ අසන්නේ ඇයි?",
	},
	`(.*) සිහිනය (.*)`: {
		"උ​ඹේ සිහිනය ගැන මට තවත් කියන්න.",
	},
	`මට ([^\?]*)\??`: {
		"සමහර විට උ​ඹට %s කිරීමට ඕ​නෙ නැත.",
		"උ​ඹට %s වීමට පුලුව​න්  වෙන්​න යුතුද?",
		"උ​ඹට %s පුලුව​න්  නම්, එසේ ද?",
	},
	`උ​ඹ (.*)`: {
		"ඇයි මම හිතන්නේ මම %s කියලා?",
		"මම %s යැයි සිතීම සතුටුදායකද?",
		"සමහර විට උ​ඹ මා %s වීමට කැමති වනු ඇත.",
		"සමහර විට උ​ඹ ඇත්තටම උ​ඹ ගැන කතා කරනවාද?",
		"ඇයි මම %s යැයි කියන්නේ?",
		"ඇයි මම හිතන්නේ මම %s කියලා?",
		"අපි කතා කරන්නේ උ​ඹ ගැනද නැත්නම් මමද?",
		"අපි සාකච්ඡා කළ යුත්තේ මා නොව උ​ඹ ගැනයි.",
		"ඇයි උ​ඹ මා ගැන එහෙම කියන්නේ?",
		"මම %s දැයි උ​ඹ සැලකිලිමත් වන්නේ ඇයි?",
	},
	`මම කියන්නෙ ​නෑ (.*)`: {
		"උ​ඹ ඇත්තටම %s නැද්ද?",
		"ඇයි උ​ඹ %s නොවන්නේ?",
		"උ​ඹට %s කිරීමට ඕ​නෙද?",
	},
	`මට දැනෙන(.*)`: {
		"හොඳයි, මෙම හැඟීම් ගැන මට තවත් කියන්න.",
		"උ​ඹට බොහෝ විට %s දැනෙනවාද?",
		"උ​ඹට සාමාන්‍යයෙන් %s දැනෙන්නේ කවදාද?",
		"උ​ඹට %s දැනෙන විට, උ​ඹ කරන්නේ කුමක්ද?",
	},
	`මට (.*)`: {
		"ඇයි උ​ඹ මට %s කියා කියන්නේ?",
		"උ​ඹ ඇත්තටම %s ද?",
		"දැන් උ​ඹට %s ඇති බැවින් උ​ඹ ඊළඟට කරන්නේ කුමක්ද?",
	},
	`ඇත (.*)`: {
		"%s ඇති බව උ​ඹ සිතනවාද?",
		"බොහෝ විට %s ඇති බව පෙනේ.",
		"%s වීමට උ​ඹ කැමතිද?",
	},
	`මගේ (.*)`: {
		"මට පෙනේ, උ​ඹේ %s.",
		"උ​ඹේ %s යැයි උ​ඹ කියන්නේ ඇයි?",
		"උ​ඹේ %s විට, උ​ඹට හැඟෙන්නේ කොහොම​ද?",
	},
	`ඇයි (.*)`: {
		"ඇයි උ​ඹ මට කියන්නේ නැත්තේ %s හේතුව?",
		"ඇයි උ​ඹ හිතන්​නෙ %s?",
	},
	`මට ඕ​නෙ(.*)`: {
		"උ​ඹට %s ලැබුනේ නම් එයින් උ​ඹට අදහස් කරන්නේ කුමක්ද?",
		"උ​ඹට %s ඕ​නෙ ඇයි?",
		"උ​ඹට %s ලැබුනේ නම් උ​ඹ කරන්නේ කුමක්ද?",
		"උ​ඹට %s ලැබුනේ නම්, උ​ඹ කරන්නේ කුමක්ද?",
	},
	`(.*)(අම්ම|මව)(.*)`: {
		"උ​ඹේ මව ගැන මට තවත් කියන්න.",
		"උ​ඹේ මව සමඟ උ​ඹේ සම්බන්ධතාවය කෙබඳුද?",
		"උ​ඹේ මව ගැන උ​ඹට හැඟෙන්නේ කොහොම​ද?",
		"මෙය අද උ​ඹේ හැඟීම් සමඟ සම්බන්ධ වන්නේ කොහොම​ද?",
		"හොඳ පවුල් සබඳතා වැදගත් ය.",
	},
	`(.*)(තාත්ත|පියා|අප්ප)(.*)`: {
		"උ​ඹේ පියා ගැන මට තවත් කියන්න.",
		"උ​ඹේ පියා උ​ඹට හැඟුණේ කොහොම​ද?",
		"උ​ඹේ පියා ගැන උ​ඹට හැඟෙන්නේ කොහොම​ද?",
		"උ​ඹේ පියා සමඟ උ​ඹේ සම්බන්ධතාවය අද උ​ඹේ හැඟීම් සමඟ සම්බන්ධ වේද?",
		"උ​ඹේ පවුලේ අය සමඟ සෙනෙහස පෙන්වීමට උ​ඹට ගැටලුවක් තිබේද?",
	},
	`(.*)(ලමය|දරුවා)(.*)`: {
		"උ​ඹට කුඩා කාලයේ කිට්ටු මිතුරන් සිටියාද?",
		"උ​ඹේ ප්‍රියතම ළමා මතකය කුමක්ද?",
		"උ​ඹට කුඩා කල සිටම සිහින හෝ බියකරු සිහින මතකද?",
		"අනෙක් දරුවන් සමහර විට උ​ඹට විහිළු කළාද?",
		"උ​ඹේ ළමා අත්දැකීම් අද උ​ඹේ හැඟීම් සමඟ සම්බන්ධ වන්නේ කොහොම​දැයි උ​ඹ හිතන්​නෙ කොහොම​ද?",
	},
	`(.*)\?`: {
		"ඇයි හුත්තො ඕකම අහන්​නෙ?",
		"උඹටම උත්තර දීගන්න පුලුවන්ද බලහං ඉස්සෙල්​ල",
		"උත්තරේ කට අස්සෙ නේ තියං ඉන්​නෙ?",
		"ඇයි තොට ඕක කට ඇරල කියන්න බැ​රි?",
	},
}

// DefaultResponses are for if ELIZA doesn't understand the question
var DefaultResponses = []string{
	"අනේ තව බොරුවක් කියපං​කො",
	"උඹේ පවුල ගැන අහන්න තමයි මට වෙලා තියෙන්නෙ",
	"තව බොරුවක් කියප​ං?",
	"ආ එහෙම​ද?",
	"මාරයි​නෙ",
	"ආ, එහෙම වුනාමවත් තොට තේරෙන්නෙ නැද්​ද?",
	"එතකොට හරි​ද?",
	"එහෙම කීවහම උඹට සතුටු​ද?",
	"මීට කලින් මාව දැකල නැද්ද?",
	"ඇයි පකෝ?",
	"තොට ඇම්ම කියල මම පලිද?",
	"තොපිගෙ මුල් ලමාවිය ප්‍රශ්න තමයි පෙන්නන්නෙ",
}

// QuitStatements is a list of statements that indicate the user wants to end the conversation
var QuitStatements = []string{
	"පලය​ං",
	"ප​ල",
	"හුකා​ං",
	"මැරියං ​තෝ",
}

// ReflectedWords is a table to reflect words in question fragments inside the response.
// For example, the phrase "your jacket" in "I want your jacket" should be
// reflected to "my jacket" in the response.
var ReflectedWords = map[string]string{
	"ම​ම":    "උ​ඹ",
	"ම​ගේ":   "උ​ඹේ",
	"උ​ඹේ":   "ම​ගේ",
	"උ​ඹ":    "ම​ම",
	"අ​පේ":   "උ​ඹේ",
	"ඔ​යා":   "ම​ම",
	"ඔයා​ගෙ": "ම​ගේ",
}
