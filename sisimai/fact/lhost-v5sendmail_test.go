// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _     __     ______                     _                 _ _ 
// |_   _|__  ___| |_   / / | |__   ___  ___| |_   \ \   / / ___| ___  ___ _ __   __| |_ __ ___   __ _(_) |
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|___\ \ / /|___ \/ __|/ _ \ '_ \ / _` | '_ ` _ \ / _` | | |
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____\ V /  ___) \__ \  __/ | | | (_| | | | | | | (_| | | |
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|     \_/  |____/|___/\___|_| |_|\__,_|_| |_| |_|\__,_|_|_|
import "testing"

func TestLhostV5sendmail(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "4.0.947", "421", "expired",         false, ""}},
		{{"02",   1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"03",   1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"04",   1, "5.0.912", "550", "hostunknown",      true, ""},
		 {"04",   2, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"05",   1, "5.0.971", "550", "blocked",         false, ""},
		 {"05",   2, "5.0.912", "550", "hostunknown",      true, ""},
		 {"05",   3, "5.0.912", "550", "hostunknown",      true, ""},
		 {"05",   4, "5.0.911", "550", "userunknown",      true, ""}},
		{{"06",   1, "5.0.909", "550", "norelaying",      false, ""}},
		{{"07",   1, "5.0.971", "554", "blocked",         false, ""},
		 {"07",   2, "5.0.912", "550", "hostunknown",      true, ""},
		 {"07",   3, "5.0.911", "550", "userunknown",      true, ""}},
	}; EngineTest(t, "V5sendmail", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1002", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1002", 2, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1003", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1004", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1005", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1005", 2, "5.0.911", "550", "userunknown",      true, ""},
		 {"1005", 3, "5.0.911", "550", "userunknown",      true, ""},
		 {"1005", 4, "5.0.911", "550", "userunknown",      true, ""},
		 {"1005", 5, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1006", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1007", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1007", 2, "5.0.911", "550", "userunknown",      true, ""},
		 {"1007", 3, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1008", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1009", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1009", 2, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1010", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1011", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1012", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1013", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1014", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1015", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1016", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1017", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1018", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1018", 2, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1019", 1, "5.0.910", "550", "filtered",        false, ""}},
		{{"1020", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1021", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1022", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1022", 2, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1022", 3, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1022", 4, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1022", 5, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1023", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1024", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1025", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1026", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1026", 2, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1026", 3, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1026", 4, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1026", 5, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1027", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1028", 1, "5.0.973", "550", "requireptr",      false, ""},
		 {"1028", 2, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1029", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1029", 2, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1030", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1031", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1032", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1033", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1033", 2, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1033", 3, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1033", 4, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1034", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1035", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1036", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1036", 2, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1037", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1038", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1038", 2, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1039", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1040", 1, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1040", 2, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1041", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1041", 2, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1042", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1043", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1043", 2, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1044", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1045", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1045", 2, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1045", 3, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1046", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1047", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1048", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1049", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1050", 1, "5.0.971", "553", "blocked",         false, ""}},
		{{"1051", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1051", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1052", 1, "5.0.971", "550", "blocked",         false, ""},
		 {"1052", 2, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1052", 3, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1052", 4, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1053", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1054", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1054", 2, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1054", 3, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1055", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1055", 2, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1055", 3, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1055", 4, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1055", 5, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1056", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1056", 2, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1057", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1058", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1059", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1060", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1061", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1062", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1064", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1065", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1066", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1066", 2, "5.7.1",   "550", "norelaying",      false, ""}},
		{{"1067", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1068", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1069", 1, "5.0.910", "550", "filtered",        false, ""}},
		{{"1071", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1072", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1072", 2, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1073", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1073", 2, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1073", 3, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1073", 4, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1073", 5, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1073", 6, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1073", 7, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1074", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1075", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1076", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1076", 2, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1076", 3, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1077", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1078", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1079", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1080", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1081", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1081", 2, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1081", 3, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1081", 4, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1081", 5, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1081", 6, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1081", 7, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1081", 8, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1081", 9, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1081",10, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1081",11, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1082", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1083", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1084", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1085", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1086", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1087", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1087", 2, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1087", 3, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1088", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1089", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1090", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1091", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1092", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1093", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1094", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1094", 2, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1095", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1095", 2, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1095", 3, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1096", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1096", 2, "5.0.911", "550", "userunknown",      true, ""},
		 {"1096", 3, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1096", 4, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1097", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1097", 2, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1098", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1100", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1101", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1102", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1103", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1104", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1105", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1106", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1106", 2, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1107", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1108", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1109", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1110", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1111", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1112", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1113", 1, "4.0.971", "421", "blocked",         false, ""}},
		{{"1114", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1115", 1, "4.0.944", "421", "networkerror",    false, ""}},
		{{"1116", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1117", 1, "4.0.973", "421", "requireptr",      false, ""}},
		{{"1118", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1118", 2, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1118", 3, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1119", 1, "4.0.947", "421", "expired",         false, ""}},
		{{"1120", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1121", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1122", 1, "4.0.971", "421", "blocked",         false, ""}},
		{{"1123", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1124", 1, "4.0.947", "421", "expired",         false, ""}},
		{{"1125", 1, "4.0.947", "421", "expired",         false, ""}},
		{{"1126", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1127", 1, "4.0.947", "421", "expired",         false, ""}},
		{{"1128", 1, "5.0.909", "550", "norelaying",      false, ""}},
		{{"1129", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1130", 1, "4.0.947", "421", "expired",         false, ""}},
		{{"1131", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1132", 1, "5.0.910", "550", "filtered",        false, ""}},
		{{"1133", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1134", 1, "4.0.947", "421", "expired",         false, ""}},
		{{"1135", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1136", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1137", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1137", 2, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1137", 3, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1138", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1139", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1139", 2, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1140", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1140", 2, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1141", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1142", 1, "5.0.911", "550", "userunknown",      true, ""}},
		{{"1143", 1, "5.0.912", "550", "hostunknown",      true, ""}},
		{{"1144", 1, "5.0.911", "550", "userunknown",      true, ""},
		 {"1144", 2, "5.0.912", "550", "hostunknown",      true, ""},
		 {"1144", 3, "5.0.911", "550", "userunknown",      true, ""}},
	}; EngineTest(t, "V5sendmail", secretlist, false)
}

