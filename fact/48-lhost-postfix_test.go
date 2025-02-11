// Copyright (C) 2025 azumakuniyuki and sisimai development team, All rights reserved.
// This software is distributed under The BSD 2-Clause License.
package fact

//  _____         _      ___ _               _        ____           _    __ _      
// |_   _|__  ___| |_   / / | |__   ___  ___| |_     |  _ \ ___  ___| |_ / _(_)_  __
//   | |/ _ \/ __| __| / /| | '_ \ / _ \/ __| __|____| |_) / _ \/ __| __| |_| \ \/ /
//   | |  __/\__ \ |_ / / | | | | | (_) \__ \ ||_____|  __/ (_) \__ \ |_|  _| |>  < 
//   |_|\___||___/\__/_/  |_|_| |_|\___/|___/\__|    |_|   \___/|___/\__|_| |_/_/\_\
import "testing"

func TestLhostPostfix(t *testing.T) {
	publiclist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"01",   1, "5.1.1",   "",    "mailererror",     false, ""}},
		{{"02",   1, "5.2.1",   "550", "userunknown",      true, ""},
		 {"02",   2, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"03",   1, "5.0.0",   "550", "filtered",        false, ""}},
		{{"04",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"05",   1, "4.1.1",   "450", "userunknown",      true, ""}},
		{{"06",   1, "5.4.4",   "",    "hostunknown",      true, ""}},
		{{"07",   1, "5.0.910", "550", "filtered",        false, ""}},
		{{"08",   1, "4.4.1",   "",    "expired",         false, ""}},
		{{"09",   1, "4.3.2",   "452", "toomanyconn",     false, ""}},
		{{"10",   1, "5.1.8",   "553", "rejected",        false, ""}},
		{{"11",   1, "5.1.8",   "553", "rejected",        false, ""},
		 {"11",   2, "5.1.8",   "553", "rejected",        false, ""}},
		{{"13",   1, "5.2.1",   "550", "userunknown",      true, ""},
		 {"13",   2, "5.2.2",   "550", "mailboxfull",     false, ""}},
		{{"14",   1, "5.1.1",   "",    "userunknown",      true, ""}},
		{{"15",   1, "4.4.1",   "",    "expired",         false, ""}},
		{{"16",   1, "5.1.6",   "550", "hasmoved",         true, ""}},
		{{"17",   1, "5.4.4",   "",    "networkerror",    false, ""}},
		{{"28",   1, "5.7.1",   "550", "notcompliantrfc", false, ""}},
		{{"29",   1, "5.7.1",   "550", "notcompliantrfc", false, ""}},
		{{"30",   1, "5.4.1",   "550", "userunknown",      true, ""}},
		{{"31",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"32",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"33",   1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"34",   1, "5.0.944", "",    "networkerror",    false, ""}},
		{{"35",   1, "5.0.0",   "550", "filtered",        false, ""}},
		{{"36",   1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"37",   1, "4.4.1",   "",    "expired",         false, ""}},
		{{"38",   1, "4.0.0",   "",    "blocked",         false, ""}},
		{{"39",   1, "5.6.0",   "554", "spamdetected",    false, ""}},
		{{"40",   1, "4.0.0",   "451", "systemerror",     false, ""}},
		{{"41",   1, "5.0.0",   "550", "policyviolation", false, ""}},
		{{"42",   1, "5.0.0",   "550", "policyviolation", false, ""}},
		{{"43",   1, "4.3.0",   "",    "mailererror",     false, ""}},
		{{"44",   1, "5.7.1",   "501", "norelaying",      false, ""}},
		{{"45",   1, "4.3.0",   "",    "mailboxfull",     false, ""}},
		{{"46",   1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"47",   1, "5.0.0",   "554", "systemerror",     false, ""}},
		{{"48",   1, "5.0.0",   "552", "toomanyconn",     false, ""}},
		{{"49",   1, "4.0.0",   "421", "blocked",         false, ""}},
		{{"50",   1, "4.0.0",   "421", "blocked",         false, ""}},
		{{"51",   1, "5.7.0",   "550", "policyviolation", false, ""}},
		{{"52",   1, "5.0.0",   "554", "suspend",         false, ""}},
		{{"53",   1, "5.0.0",   "504", "syntaxerror",     false, ""}},
		{{"54",   1, "5.7.1",   "550", "rejected",        false, ""}},
		{{"55",   1, "5.0.0",   "552", "toomanyconn",     false, ""}},
		{{"56",   1, "4.4.2",   "",    "networkerror",    false, ""}},
		{{"57",   1, "5.2.1",   "550", "userunknown",      true, ""}},
		{{"58",   1, "5.7.1",   "550", "badreputation",   false, ""}},
		{{"59",   1, "5.2.1",   "550", "speeding",        false, ""}},
		{{"60",   1, "4.0.0",   "",    "requireptr",      false, ""}},
		{{"61",   1, "5.0.0",   "550", "suspend",         false, ""}},
		{{"62",   1, "5.0.0",   "550", "virusdetected",   false, ""}},
		{{"63",   1, "5.2.2",   "552", "mailboxfull",     false, ""}},
		{{"64",   1, "5.0.900", "",    "undefined",       false, ""}},
		{{"65",   1, "5.0.0",   "550", "securityerror",   false, ""}},
		{{"66",   1, "5.7.9",   "554", "policyviolation", false, ""}},
		{{"67",   1, "5.7.9",   "554", "policyviolation", false, ""}},
		{{"68",   1, "5.0.0",   "554", "policyviolation", false, ""}},
		{{"69",   1, "5.7.9",   "554", "policyviolation", false, ""}},
		{{"70",   1, "5.7.26",  "550", "authfailure",     false, ""}},
		{{"71",   1, "5.7.1",   "554", "authfailure",     false, ""}},
		{{"72",   1, "5.7.1",   "550", "authfailure",     false, ""}},
		{{"73",   1, "5.7.1",   "550", "authfailure",     false, ""}},
		{{"74",   1, "4.7.0",   "421", "rejected",        false, ""}},
		{{"75",   1, "4.3.0",   "451", "systemerror",     false, ""}},
		{{"76",   1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"77",   1, "5.0.0",   "554", "norelaying",      false, ""}},
		{{"78",   1, "5.0.0",   "554", "notcompliantrfc", false, ""}},
	}; EngineTest(t, "Postfix", publiclist, true)

	secretlist := [][]IsExpected{
		// Label, Index, Status, ReplyCode, Reason, HardBounce, AnotherOne
		{{"1001", 1, "5.0.0",   "550", "rejected",        false, ""}},
		{{"1002", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1003", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1004", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1005", 1, "5.0.0",   "554", "filtered",        false, ""}},
		{{"1006", 1, "5.7.1",   "550", "userunknown",      true, ""}},
		{{"1007", 1, "5.0.0",   "554", "filtered",        false, ""}},
		{{"1008", 1, "5.0.910", "",    "filtered",        false, ""}},
		{{"1009", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1010", 1, "5.0.0",   "",    "hostunknown",      true, ""}},
		{{"1011", 1, "5.0.0",   "551", "systemerror",     false, ""}},
		{{"1012", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1013", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1014", 1, "5.1.1",   "",    "userunknown",      true, ""}},
		{{"1015", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1016", 1, "4.3.2",   "452", "toomanyconn",     false, ""}},
		{{"1017", 1, "4.4.1",   "",    "expired",         false, ""}},
		{{"1018", 1, "5.4.6",   "",    "systemerror",     false, ""}},
		{{"1019", 1, "5.7.1",   "553", "userunknown",      true, ""}},
		{{"1020", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1021", 1, "4.4.1",   "",    "expired",         false, ""}},
		{{"1022", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1023", 1, "5.0.0",   "550", "blocked",         false, ""}},
		{{"1024", 1, "5.1.1",   "",    "userunknown",      true, ""}},
		{{"1025", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1026", 1, "4.4.1",   "",    "expired",         false, ""}},
		{{"1027", 1, "5.4.6",   "",    "systemerror",     false, ""}},
		{{"1028", 1, "5.0.0",   "551", "suspend",         false, ""}},
		{{"1029", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1030", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1031", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1032", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1033", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1034", 1, "5.0.0",   "550", "rejected",        false, ""}},
		{{"1035", 1, "4.2.2",   "",    "mailboxfull",     false, ""}},
		{{"1036", 1, "5.4.4",   "",    "hostunknown",      true, ""}},
		{{"1037", 1, "5.0.0",   "550", "rejected",        false, ""}},
		{{"1038", 1, "5.0.0",   "550", "blocked",         false, ""}},
		{{"1039", 1, "5.1.1",   "",    "userunknown",      true, ""}},
		{{"1040", 1, "5.7.1",   "550", "userunknown",      true, ""}},
		{{"1041", 1, "5.1.1",   "",    "userunknown",      true, ""}},
		{{"1042", 1, "5.4.4",   "",    "networkerror",    false, ""}},
		{{"1043", 1, "5.1.6",   "550", "hasmoved",         true, ""}},
		{{"1044", 1, "5.3.4",   "",    "mesgtoobig",      false, ""}},
		{{"1045", 1, "5.3.4",   "",    "mesgtoobig",      false, ""}},
		{{"1046", 1, "5.0.0",   "534", "mesgtoobig",      false, ""}},
		{{"1047", 1, "5.7.1",   "554", "mesgtoobig",      false, ""}},
		{{"1048", 1, "5.1.1",   "550", "userunknown",      true, ""},
		 {"1048", 2, "5.1.1",   "550", "userunknown",      true, ""},
		 {"1048", 3, "5.1.1",   "550", "userunknown",      true, ""},
		 {"1048", 4, "5.1.1",   "550", "userunknown",      true, ""},
		 {"1048", 5, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1049", 1, "5.0.0",   "550", "hostunknown",      true, ""}},
		{{"1050", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1051", 1, "5.7.1",   "553", "norelaying",      false, ""}},
		{{"1052", 1, "5.7.1",   "550", "spamdetected",    false, ""}},
		{{"1053", 1, "5.4.6",   "",    "systemerror",     false, ""}},
		{{"1054", 1, "5.1.1",   "",    "userunknown",      true, ""}},
		{{"1055", 1, "5.2.1",   "550", "filtered",        false, ""}},
		{{"1056", 1, "5.1.1",   "",    "mailererror",     false, ""}},
		{{"1057", 1, "5.2.1",   "550", "userunknown",      true, ""},
		 {"1057", 2, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1058", 1, "5.0.0",   "550", "filtered",        false, ""}},
		{{"1059", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1060", 1, "4.1.1",   "450", "userunknown",      true, ""}},
		{{"1061", 1, "5.4.4",   "",    "hostunknown",      true, ""}},
		{{"1062", 1, "5.0.910", "550", "filtered",        false, ""}},
		{{"1063", 1, "5.1.1",   "",    "mailererror",     false, ""}},
		{{"1064", 1, "5.0.0",   "",    "hostunknown",      true, ""}},
		{{"1065", 1, "5.0.0",   "",    "networkerror",    false, ""}},
		{{"1066", 1, "5.0.0",   "554", "norelaying",      false, ""}},
		{{"1067", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1068", 1, "5.0.0",   "554", "norelaying",      false, ""}},
		{{"1069", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1070", 1, "5.0.944", "",    "networkerror",    false, ""}},
		{{"1071", 1, "5.0.922", "",    "mailboxfull",     false, ""}},
		{{"1072", 1, "5.0.901", "554", "onhold",          false, ""}},
		{{"1073", 1, "4.0.0",   "452", "mailboxfull",     false, ""}},
		{{"1074", 1, "5.0.0",   "550", "mailboxfull",     false, ""}},
		{{"1075", 1, "5.7.0",   "",    "mailboxfull",     false, ""}},
		{{"1076", 1, "5.0.0",   "554", "filtered",        false, ""}},
		{{"1077", 1, "5.7.1",   "553", "norelaying",      false, ""}},
		{{"1078", 1, "5.0.0",   "550", "norelaying",      false, ""}},
		{{"1079", 1, "5.7.1",   "550", "spamdetected",    false, ""}},
		{{"1080", 1, "5.7.1",   "554", "spamdetected",    false, ""}},
		{{"1081", 1, "5.0.0",   "550", "spamdetected",    false, ""}},
		{{"1082", 1, "5.0.0",   "550", "spamdetected",    false, ""}},
		{{"1083", 1, "5.0.0",   "550", "spamdetected",    false, ""}},
		{{"1084", 1, "5.7.1",   "554", "spamdetected",    false, ""}},
		{{"1085", 1, "5.7.1",   "554", "spamdetected",    false, ""}},
		{{"1086", 1, "5.0.0",   "550", "spamdetected",    false, ""}},
		{{"1087", 1, "5.0.0",   "550", "spamdetected",    false, ""}},
		{{"1088", 1, "5.6.0",   "554", "spamdetected",    false, ""}},
		{{"1089", 1, "5.7.1",   "554", "spamdetected",    false, ""}},
		{{"1090", 1, "5.7.1",   "554", "spamdetected",    false, ""}},
		{{"1091", 1, "5.0.0",   "500", "spamdetected",    false, ""}},
		{{"1092", 1, "5.0.0",   "554", "spamdetected",    false, ""}},
		{{"1093", 1, "5.7.1",   "554", "spamdetected",    false, ""}},
		{{"1094", 1, "5.7.1",   "550", "policyviolation", false, ""}},
		{{"1095", 1, "5.0.0",   "554", "spamdetected",    false, ""}},
		{{"1096", 1, "5.0.0",   "554", "spamdetected",    false, ""}},
		{{"1097", 1, "5.7.3",   "553", "spamdetected",    false, ""}},
		{{"1098", 1, "5.7.1",   "550", "spamdetected",    false, ""}},
		{{"1099", 1, "5.7.1",   "550", "spamdetected",    false, ""}},
		{{"1100", 1, "5.0.0",   "554", "spamdetected",    false, ""}},
		{{"1101", 1, "5.0.0",   "554", "virusdetected",   false, ""}},
		{{"1102", 1, "5.7.1",   "550", "spamdetected",    false, ""}},
		{{"1103", 1, "5.0.0",   "550", "spamdetected",    false, ""}},
		{{"1104", 1, "5.0.0",   "550", "spamdetected",    false, ""}},
		{{"1105", 1, "5.0.0",   "551", "spamdetected",    false, ""}},
		{{"1106", 1, "5.0.0",   "550", "spamdetected",    false, ""}},
		{{"1107", 1, "5.7.1",   "554", "spamdetected",    false, ""}},
		{{"1108", 1, "5.0.0",   "550", "spamdetected",    false, ""}},
		{{"1109", 1, "5.7.1",   "550", "spamdetected",    false, ""}},
		{{"1110", 1, "5.7.1",   "550", "spamdetected",    false, ""}},
		{{"1111", 1, "5.0.0",   "550", "spamdetected",    false, ""}},
		{{"1112", 1, "5.0.0",   "554", "spamdetected",    false, ""}},
		{{"1113", 1, "5.7.1",   "550", "spamdetected",    false, ""}},
		{{"1114", 1, "5.0.0",   "550", "spamdetected",    false, ""}},
		{{"1115", 1, "5.0.0",   "554", "blocked",         false, ""}},
		{{"1116", 1, "5.0.0",   "550", "spamdetected",    false, ""}},
		{{"1117", 1, "5.0.0",   "550", "spamdetected",    false, ""}},
		{{"1118", 1, "5.0.0",   "554", "spamdetected",    false, ""}},
		{{"1119", 1, "5.0.0",   "553", "spamdetected",    false, ""}},
		{{"1120", 1, "5.7.1",   "550", "spamdetected",    false, ""}},
		{{"1121", 1, "5.3.0",   "554", "spamdetected",    false, ""}},
		{{"1122", 1, "5.4.4",   "",    "hostunknown",      true, ""}},
		{{"1123", 1, "5.7.1",   "554", "userunknown",      true, ""}},
		{{"1124", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1125", 1, "5.2.3",   "",    "mailboxfull",     false, ""}},
		{{"1126", 1, "5.0.0",   "",    "systemerror",     false, ""}},
		{{"1127", 1, "5.7.17",  "550", "userunknown",      true, ""}},
		{{"1128", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1129", 1, "5.0.0",   "554", "filtered",        false, ""}},
		{{"1130", 1, "5.0.0",   "552", "mailboxfull",     false, ""}},
		{{"1131", 1, "5.2.3",   "",    "mailboxfull",     false, ""}},
		{{"1132", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1133", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1134", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1135", 1, "5.2.1",   "550", "suspend",         false, ""}},
		{{"1136", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1137", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1138", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1139", 1, "5.1.3",   "501", "userunknown",      true, ""}},
		{{"1140", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1141", 1, "5.0.0",   "",    "filtered",        false, ""}},
		{{"1142", 1, "5.0.0",   "550", "blocked",         false, ""}},
		{{"1143", 1, "5.3.0",   "553", "userunknown",      true, ""}},
		{{"1144", 1, "5.0.0",   "554", "suspend",         false, ""}},
		{{"1145", 1, "5.0.0",   "550", "rejected",        false, ""}},
		{{"1146", 1, "5.1.3",   "",    "userunknown",      true, ""}},
		{{"1147", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1148", 1, "5.2.1",   "550", "userunknown",      true, ""}},
		{{"1149", 1, "5.2.2",   "550", "mailboxfull",     false, ""}},
		{{"1150", 1, "5.0.910", "",    "filtered",        false, ""}},
		{{"1151", 1, "5.0.0",   "550", "spamdetected",    false, ""}},
		{{"1152", 1, "5.3.0",   "553", "blocked",         false, ""}},
		{{"1153", 1, "5.7.1",   "550", "badreputation",   false, ""}},
		{{"1154", 1, "4.7.0",   "421", "blocked",         false, ""}},
		{{"1155", 1, "5.1.0",   "550", "userunknown",      true, ""}},
		{{"1156", 1, "5.1.0",   "550", "userunknown",      true, ""}},
		{{"1157", 1, "4.0.0",   "",    "blocked",         false, ""}},
		{{"1158", 1, "5.6.0",   "554", "spamdetected",    false, ""}},
		{{"1159", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1160", 1, "4.0.0",   "451", "systemerror",     false, ""}},
		{{"1161", 1, "5.0.0",   "",    "mailboxfull",     false, ""}},
		{{"1162", 1, "5.0.0",   "550", "policyviolation", false, ""}},
		{{"1163", 1, "5.0.0",   "550", "policyviolation", false, ""}},
		{{"1164", 1, "5.0.0",   "550", "blocked",         false, ""}},
		{{"1165", 1, "5.5.0",   "550", "userunknown",      true, ""}},
		{{"1166", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1167", 1, "4.0.0",   "",    "blocked",         false, ""}},
		{{"1168", 1, "5.0.0",   "",    "rejected",        false, ""}},
		{{"1169", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1170", 1, "5.0.0",   "550", "requireptr",      false, ""}},
		{{"1171", 1, "5.2.0",   "",    "mailboxfull",     false, ""}},
		{{"1172", 1, "4.3.0",   "",    "mailererror",     false, ""}},
		{{"1173", 1, "4.4.2",   "",    "networkerror",    false, ""}},
		{{"1174", 1, "4.3.2",   "451", "notaccept",       false, ""}},
		{{"1175", 1, "5.7.9",   "554", "policyviolation", false, ""}},
		{{"1176", 1, "5.7.1",   "554", "userunknown",      true, ""}},
		{{"1177", 1, "5.7.1",   "550", "userunknown",      true, ""}},
		{{"1178", 1, "5.7.1",   "550", "blocked",         false, ""}},
		{{"1179", 1, "5.7.1",   "501", "norelaying",      false, ""}},
		{{"1180", 1, "5.4.1",   "550", "userunknown",      true, ""}},
		{{"1181", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1182", 1, "5.7.0",   "550", "spamdetected",    false, ""}},
		{{"1183", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1184", 1, "5.7.1",   "550", "norelaying",      false, ""}},
		{{"1185", 1, "4.0.0",   "451", "systemerror",     false, ""}},
		{{"1186", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1187", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1188", 1, "4.4.1",   "",    "expired",         false, ""}},
		{{"1189", 1, "5.4.4",   "",    "hostunknown",      true, ""}},
		{{"1190", 1, "5.1.1",   "",    "userunknown",      true, ""}},
		{{"1191", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1192", 1, "5.1.1",   "550", "speeding",        false, ""}},
		{{"1193", 1, "5.0.0",   "550", "filtered",        false, ""}},
		{{"1194", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1195", 1, "4.4.1",   "",    "expired",         false, ""}},
		{{"1196", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1197", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1198", 1, "5.0.0",   "554", "systemerror",     false, ""}},
		{{"1199", 1, "5.0.0",   "552", "toomanyconn",     false, ""}},
		{{"1200", 1, "4.0.0",   "421", "blocked",         false, ""}},
		{{"1201", 1, "4.0.0",   "421", "blocked",         false, ""}},
		{{"1202", 1, "5.7.0",   "550", "policyviolation", false, ""}},
		{{"1203", 1, "5.0.0",   "554", "suspend",         false, ""}},
		{{"1204", 1, "5.0.0",   "504", "syntaxerror",     false, ""}},
		{{"1205", 1, "5.7.1",   "550", "rejected",        false, ""}},
		{{"1206", 1, "5.0.0",   "552", "toomanyconn",     false, ""}},
		{{"1207", 1, "5.0.0",   "550", "toomanyconn",     false, ""}},
		{{"1208", 1, "5.0.0",   "550", "toomanyconn",     false, ""}},
		{{"1209", 1, "4.4.2",   "",    "networkerror",    false, ""}},
		{{"1210", 1, "5.0.0",   "550", "authfailure",     false, ""}},
		{{"1211", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1212", 1, "5.2.1",   "550", "userunknown",      true, ""}},
		{{"1213", 1, "5.1.1",   "550", "userunknown",      true, ""}},
		{{"1214", 1, "5.2.1",   "550", "speeding",        false, ""}},
		{{"1215", 1, "5.2.1",   "550", "speeding",        false, ""}},
		{{"1216", 1, "4.0.0",   "",    "requireptr",      false, ""}},
		{{"1217", 1, "4.0.0",   "",    "requireptr",      false, ""}},
		{{"1218", 1, "4.0.0",   "",    "requireptr",      false, ""}},
		{{"1219", 1, "5.0.0",   "550", "suspend",         false, ""}},
		{{"1220", 1, "5.0.0",   "550", "virusdetected",   false, ""}},
		{{"1221", 1, "5.1.1",   "",    "userunknown",      true, ""}},
		{{"1222", 1, "5.2.2",   "552", "mailboxfull",     false, ""}},
		{{"1223", 1, "5.7.9",   "554", "policyviolation", false, ""}},
		{{"1224", 1, "5.7.9",   "554", "policyviolation", false, ""}},
		{{"1225", 1, "5.0.0",   "554", "policyviolation", false, ""}},
		{{"1226", 1, "5.7.9",   "554", "policyviolation", false, ""}},
		{{"1227", 1, "5.7.26",  "550", "authfailure",     false, ""}},
		{{"1228", 1, "5.7.1",   "554", "authfailure",     false, ""}},
		{{"1229", 1, "5.7.1",   "550", "authfailure",     false, ""}},
		{{"1230", 1, "5.7.1",   "550", "authfailure",     false, ""}},
		{{"1231", 1, "5.7.9",   "550", "policyviolation", false, ""},
		 {"1231", 2, "5.7.1",   "550", "authfailure",     false, ""},
		 {"1231", 3, "5.7.1",   "550", "authfailure",     false, ""}},
		{{"1232", 1, "4.7.0",   "421", "rejected",        false, ""}},
		{{"1233", 1, "5.0.0",   "550", "blocked",         false, ""}},
		{{"1234", 1, "5.0.0",   "553", "rejected",        false, ""}},
		{{"1235", 1, "5.0.0",   "554", "spamdetected",    false, ""}},
		{{"1236", 1, "5.0.0",   "550", "badreputation",   false, ""}},
		{{"1237", 1, "5.0.0",   "550", "norelaying",      false, ""}},
		{{"1238", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1239", 1, "5.0.0",   "550", "blocked",         false, ""}},
		{{"1240", 1, "5.0.0",   "550", "rejected",        false, ""}},
		{{"1241", 1, "5.0.0",   "550", "rejected",        false, ""}},
		{{"1242", 1, "5.0.0",   "550", "spamdetected",    false, ""}},
		{{"1243", 1, "5.0.0",   "554", "badreputation",   false, ""}},
		{{"1244", 1, "5.0.972", "550", "policyviolation", false, ""}}, // 5.8.5 is an invalid status
		{{"1245", 1, "5.0.0",   "554", "blocked",         false, ""}},
		{{"1246", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1247", 1, "5.0.0",   "550", "norelaying",      false, ""}},
		{{"1248", 1, "5.0.0",   "550", "blocked",         false, ""}},
		{{"1249", 1, "5.0.0",   "550", "blocked",         false, ""}},
		{{"1250", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1251", 1, "5.0.0",   "550", "spamdetected",    false, ""}},
		{{"1252", 1, "5.0.0",   "",    "onhold",          false, ""}},
		{{"1253", 1, "5.0.0",   "554", "spamdetected",    false, ""}},
		{{"1254", 1, "5.0.0",   "554", "policyviolation", false, ""}},
		{{"1255", 1, "5.4.6",   "554", "systemerror",     false, ""}},
		{{"1256", 1, "5.5.1",   "554", "blocked",         false, ""}},
		{{"1257", 1, "5.0.0",   "550", "notaccept",        true, ""}},
		{{"1258", 1, "5.0.0",   "550", "rejected",        false, ""}},
		{{"1259", 1, "5.0.0",   "",    "onhold",          false, ""}},
		{{"1260", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1261", 1, "5.0.0",   "550", "norelaying",      false, ""}},
		{{"1262", 1, "5.0.0",   "550", "norelaying",      false, ""}},
		{{"1263", 1, "5.0.0",   "550", "filtered",        false, ""}},
		{{"1264", 1, "5.0.0",   "550", "userunknown",      true, ""}},
		{{"1265", 1, "5.0.0",   "554", "rejected",        false, ""}},
		{{"1266", 1, "5.0.0",   "550", "suspend",         false, ""}},
		{{"1267", 1, "5.0.0",   "550", "onhold",          false, ""}}, // spamdetected
		{{"1268", 1, "5.0.0",   "550", "suspend",         false, ""}},
		{{"1269", 1, "5.0.0",   "550", "virusdetected",   false, ""}},
		{{"1270", 1, "5.0.0",   "554", "norelaying",      false, ""}},
		{{"1271", 1, "5.0.0",   "554", "notcompliantrfc", false, ""}},
		{{"1272", 1, "5.0.0",   "550", "rejected",        false, ""}},
		{{"1273", 1, "5.0.0",   "550", "rejected",        false, ""}},
		{{"1274", 1, "5.0.939", "",    "mailererror",     false, ""}},
		{{"1275", 1, "5.4.14",  "554", "networkerror",    false, ""},
		 {"1275", 2, "5.4.14",  "554", "networkerror",    false, ""}},
		{{"1276", 1, "5.7.26",  "550", "authfailure",     false, ""}},
		{{"1277", 1, "5.7.26",  "550", "authfailure",     false, ""}},
		{{"1278", 1, "5.7.25",  "550", "requireptr",      false, ""}},
		{{"1279", 1, "5.2.2",   "552", "mailboxfull",     false, ""}},
		{{"1280", 1, "5.7.1",   "550", "notcompliantrfc", false, ""}},
	}; EngineTest(t, "Postfix", secretlist, false)
}

