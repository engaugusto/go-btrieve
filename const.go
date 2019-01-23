// Copyright 2019 Barchart.com, Inc. All rights reserved.
// Use of this source code is governed by a GNU-style
// license that can be found in the LICENSE file.

/*
	Package btrieve contains constants and function that
	help in accessing Btrieve databases.
*/

package btrieve

// Openation Codes
type OP_CODE uint16

const (
	B_OPEN              OP_CODE = 0
	B_CLOSE             OP_CODE = 1
	B_INSERT            OP_CODE = 2
	B_UPDATE            OP_CODE = 3
	B_DELETE            OP_CODE = 4
	B_GET_EQUAL         OP_CODE = 5
	B_GET_NEXT          OP_CODE = 6
	B_GET_PREVIOUS      OP_CODE = 7
	B_GET_GT            OP_CODE = 8
	B_GET_GE            OP_CODE = 9
	B_GET_LT            OP_CODE = 10
	B_GET_LE            OP_CODE = 11
	B_GET_FIRST         OP_CODE = 12
	B_GET_LAST          OP_CODE = 13
	B_CREATE            OP_CODE = 14
	B_STAT              OP_CODE = 15
	B_EXTEND            OP_CODE = 16
	B_SET_DIR           OP_CODE = 17
	B_GET_DIR           OP_CODE = 18
	B_BEGIN_TRAN        OP_CODE = 19
	B_END_TRAN          OP_CODE = 20
	B_ABORT_TRAN        OP_CODE = 21
	B_GET_POSITION      OP_CODE = 22
	B_GET_DIRECT        OP_CODE = 23
	B_STEP_NEXT         OP_CODE = 24
	B_STOP              OP_CODE = 25
	B_VERSION           OP_CODE = 26
	B_UNLOCK            OP_CODE = 27
	B_RESET             OP_CODE = 28
	B_SET_OWNER         OP_CODE = 29
	B_CLEAR_OWNER       OP_CODE = 30
	B_BUILD_INDEX       OP_CODE = 31
	B_DROP_INDEX        OP_CODE = 32
	B_STEP_FIRST        OP_CODE = 33
	B_STEP_LAST         OP_CODE = 34
	B_STEP_PREVIOUS     OP_CODE = 35
	B_GET_NEXT_EXTENDED OP_CODE = 36
	B_GET_PREV_EXTENDED OP_CODE = 37
	B_STEP_NEXT_EXT     OP_CODE = 38
	B_STEP_PREVIOUS_EXT OP_CODE = 39
	B_EXT_INSERT        OP_CODE = 40
	B_MISC_DATA         OP_CODE = 41
	B_CONTINUOUS        OP_CODE = 42
	B_SEEK_PERCENT      OP_CODE = 44
	B_GET_PERCENT       OP_CODE = 45
	B_CHUNK_UPDATE      OP_CODE = 53
)

type RET_CODE uint16

const (
	B_NO_ERROR                       RET_CODE = 0
	B_INVALID_FUNCTION               RET_CODE = 1
	B_IO_ERROR                       RET_CODE = 2
	B_FILE_NOT_OPEN                  RET_CODE = 3
	B_KEY_VALUE_NOT_FOUND            RET_CODE = 4
	B_DUPLICATE_KEY_VALUE            RET_CODE = 5
	B_INVALID_KEYNUMBER              RET_CODE = 6
	B_DIFFERENT_KEYNUMBER            RET_CODE = 7
	B_POSITION_NOT_SET               RET_CODE = 8
	B_END_OF_FILE                    RET_CODE = 9
	B_MODIFIABLE_KEYVALUE_ERROR      RET_CODE = 10
	B_FILENAME_BAD                   RET_CODE = 11
	B_FILE_NOT_FOUND                 RET_CODE = 12
	B_EXTENDED_FILE_ERROR            RET_CODE = 13
	B_PREIMAGE_OPEN_ERROR            RET_CODE = 14
	B_PREIMAGE_IO_ERROR              RET_CODE = 15
	B_EXPANSION_ERROR                RET_CODE = 16
	B_CLOSE_ERROR                    RET_CODE = 17
	B_DISKFULL                       RET_CODE = 18
	B_UNRECOVERABLE_ERROR            RET_CODE = 19
	B_RECORD_MANAGER_INACTIVE        RET_CODE = 20
	B_KEYBUFFER_TOO_SHORT            RET_CODE = 21
	B_DATALENGTH_ERROR               RET_CODE = 22
	B_POSITIONBLOCK_LENGTH           RET_CODE = 23
	B_PAGE_SIZE_ERROR                RET_CODE = 24
	B_CREATE_IO_ERROR                RET_CODE = 25
	B_NUMBER_OF_KEYS                 RET_CODE = 26
	B_INVALID_KEY_POSITION           RET_CODE = 27
	B_INVALID_RECORD_LENGTH          RET_CODE = 28
	B_INVALID_KEYLENGTH              RET_CODE = 29
	B_NOT_A_BTRIEVE_FILE             RET_CODE = 30
	B_FILE_ALREADY_EXTENDED          RET_CODE = 31
	B_EXTEND_IO_ERROR                RET_CODE = 32
	B_BTR_CANNOT_UNLOAD              RET_CODE = 33
	B_INVALID_EXTENSION_NAME         RET_CODE = 34
	B_DIRECTORY_ERROR                RET_CODE = 35
	B_TRANSACTION_ERROR              RET_CODE = 36
	B_TRANSACTION_IS_ACTIVE          RET_CODE = 37
	B_TRANSACTION_FILE_IO_ERROR      RET_CODE = 38
	B_END_TRANSACTION_ERROR          RET_CODE = 39
	B_TRANSACTION_MAX_FILES          RET_CODE = 40
	B_OPERATION_NOT_ALLOWED          RET_CODE = 41
	B_INCOMPLETE_ACCEL_ACCESS        RET_CODE = 42
	B_INVALID_RECORD_ADDRESS         RET_CODE = 43
	B_NULL_KEYPATH                   RET_CODE = 44
	B_INCONSISTENT_KEY_FLAGS         RET_CODE = 45
	B_ACCESS_TO_FILE_DENIED          RET_CODE = 46
	B_MAXIMUM_OPEN_FILES             RET_CODE = 47
	B_INVALID_ALT_SEQUENCE_DEF       RET_CODE = 48
	B_KEY_TYPE_ERROR                 RET_CODE = 49
	B_OWNER_ALREADY_SET              RET_CODE = 50
	B_INVALID_OWNER                  RET_CODE = 51
	B_ERROR_WRITING_CACHE            RET_CODE = 52
	B_INVALID_INTERFACE              RET_CODE = 53
	B_VARIABLE_PAGE_ERROR            RET_CODE = 54
	B_AUTOINCREMENT_ERROR            RET_CODE = 55
	B_INCOMPLETE_INDEX               RET_CODE = 56
	B_EXPANED_MEM_ERROR              RET_CODE = 57
	B_COMPRESS_BUFFER_TOO_SHORT      RET_CODE = 58
	B_FILE_ALREADY_EXISTS            RET_CODE = 59
	B_REJECT_COUNT_REACHED           RET_CODE = 60
	B_SMALL_EX_GET_BUFFER_ERROR      RET_CODE = 61
	B_INVALID_GET_EXPRESSION         RET_CODE = 62
	B_INVALID_EXT_INSERT_BUFF        RET_CODE = 63
	B_OPTIMIZE_LIMIT_REACHED         RET_CODE = 64
	B_INVALID_EXTRACTOR              RET_CODE = 65
	B_RI_TOO_MANY_DATABASES          RET_CODE = 66
	B_RIDDF_CANNOT_OPEN              RET_CODE = 67
	B_RI_CASCADE_TOO_DEEP            RET_CODE = 68
	B_RI_CASCADE_ERROR               RET_CODE = 69
	B_RI_VIOLATION                   RET_CODE = 71
	B_RI_REFERENCED_FILE_CANNOT_OPEN RET_CODE = 72
	B_RI_OUT_OF_SYNC                 RET_CODE = 73
	B_END_CHANGED_TO_ABORT           RET_CODE = 74
	B_RI_CONFLICT                    RET_CODE = 76
	B_CANT_LOOP_IN_SERVER            RET_CODE = 77
	B_DEAD_LOCK                      RET_CODE = 78
	B_PROGRAMMING_ERROR              RET_CODE = 79
	B_CONFLICT                       RET_CODE = 80
	B_LOCKERROR                      RET_CODE = 81
	B_LOST_POSITION                  RET_CODE = 82
	B_READ_OUTSIDE_TRANSACTION       RET_CODE = 83
	B_RECORD_INUSE                   RET_CODE = 84
	B_FILE_INUSE                     RET_CODE = 85
	B_FILE_TABLE_FULL                RET_CODE = 86
	B_NOHANDLES_AVAILABLE            RET_CODE = 87
	B_INCOMPATIBLE_MODE_ERROR        RET_CODE = 88
	B_DEVICE_TABLE_FULL              RET_CODE = 90
	B_SERVER_ERROR                   RET_CODE = 91
	B_TRANSACTION_TABLE_FULL         RET_CODE = 92
	B_INCOMPATIBLE_LOCK_TYPE         RET_CODE = 93
	B_PERMISSION_ERROR               RET_CODE = 94
	B_SESSION_NO_LONGER_VALID        RET_CODE = 95
	B_COMMUNICATIONS_ERROR           RET_CODE = 96
	B_DATA_MESSAGE_TOO_SMALL         RET_CODE = 97
	B_INTERNAL_TRANSACTION_ERROR     RET_CODE = 98
	B_REQUESTER_CANT_ACCESS_RUNTIME  RET_CODE = 99
	B_NO_CACHE_BUFFERS_AVAIL         RET_CODE = 100
	B_NO_OS_MEMORY_AVAIL             RET_CODE = 101
	B_NO_STACK_AVAIL                 RET_CODE = 102
	B_CHUNK_OFFSET_TOO_LONG          RET_CODE = 103
	B_LOCALE_ERROR                   RET_CODE = 104
	B_CANNOT_CREATE_WITH_BAT         RET_CODE = 105
	B_CHUNK_CANNOT_GET_NEXT          RET_CODE = 106
	B_CHUNK_INCOMPATIBLE_FILE        RET_CODE = 107
	B_TRANSACTION_TOO_COMPLEX        RET_CODE = 109
	B_NO_SYSTEM_LOCKS_AVAILABLE      RET_CODE = 130
	B_FILE_FULL                      RET_CODE = 132
	B_MORE_THAN_5_CONCURRENT_USERS   RET_CODE = 133
	B_ISR_NOT_FOUND                  RET_CODE = 134
	B_ISR_FORMAT_INVALID             RET_CODE = 135
	B_CANNOT_CONVERT_RP              RET_CODE = 137
	B_INVALID_NULL_INDICATOR         RET_CODE = 138
	B_INVALID_KEY_OPTION             RET_CODE = 139
	B_INCOMPATIBLE_CLOSE             RET_CODE = 140
	B_INVALID_USERNAME               RET_CODE = 141
	B_INVALID_DATABASE               RET_CODE = 142
	B_NO_SSQL_RIGHTS                 RET_CODE = 143
	B_ALREADY_LOGGED_IN              RET_CODE = 144
	B_NO_DATABASE_SERVICES           RET_CODE = 145
	B_DUPLICATE_SYSTEM_KEY           RET_CODE = 146
	B_LOG_SEGMENT_MISSING            RET_CODE = 147
	B_ROLL_FORWARD_ERROR             RET_CODE = 148
	B_SYSTEM_KEY_INTERNAL            RET_CODE = 149
	B_DBS_INTERNAL_ERROR             RET_CODE = 150
	B_NESTING_DEPTH_ERROR            RET_CODE = 151
)

const (
	MAX_KEY_SIZE = uint8(255)
)
