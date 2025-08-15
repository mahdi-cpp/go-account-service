import "time"

type User struct {
// Core Identity & Basic Information
ID               string    `json:"id"`                // شناسه یکتای کاربر
Username         string    `json:"username"`          // نام کاربری (برای لاگین و شناسایی یکتا)
DisplayName      string    `json:"displayName"`       // نام نمایشی (چیزی که دیگران می‌بینند، می‌تواند متفاوت از Username باشد)
PhoneNumber      string    `json:"phoneNumber"`       // شماره تلفن کاربر
Email            string    `json:"email"`             // ایمیل کاربر
FirstName        string    `json:"firstName"`         // نام
LastName         string    `json:"lastName"`          // نام خانوادگی
Bio              string    `json:"bio"`               // بیوگرافی کوتاه یا "درباره من"
AvatarURL        string    `json:"avatarURL"`         // URL تصویر پروفایل کاربر
IsVerified       bool      `json:"isVerified"`        // نشان می‌دهد که آیا حساب کاربری تایید شده است یا خیر (مثلاً برای حساب‌های رسمی)

	// Presence & Connectivity (for chat, social apps)
	IsOnline         bool      `json:"isOnline"`          // وضعیت آنلاین بودن کاربر
	LastSeen         time.Time `json:"lastSeen"`          // آخرین باری که کاربر آنلاین دیده شده است
	StatusMessage    string    `json:"statusMessage"`     // پیام وضعیت کاربر (مثلاً "مشغول", "در دسترس")

	// Privacy & Social Features (for social, photo, music apps)
	ProfileVisibility string    `json:"profileVisibility"` // وضعیت نمایش پروفایل: "public", "private", "friendsOnly"
	FollowerCount    int       `json:"followerCount"`     // تعداد دنبال‌کننده‌ها
	FollowingCount   int       `json:"followingCount"`    // تعداد دنبال‌شوندگان
	BlockedUserIDs   []string  `json:"blockedUserIDs"`    // لیست شناسه‌های کاربرانی که توسط این کاربر بلاک شده‌اند

	// Preferences & Customization
	PreferredLanguage string    `json:"preferredLanguage"` // زبان ترجیحی کاربر (مثلاً "en-US", "fa-IR")
	Timezone         string    `json:"timezone"`          // منطقه زمانی کاربر (مثلاً "Asia/Tehran")
	ThemePreference  string    `json:"themePreference"`   // ترجیح تم رابط کاربری: "light", "dark", "system"
	Interests        []string  `json:"interests"`         // لیستی از علایق کاربر (برای پیشنهاد محتوا در موسیقی، عکس، آموزش)

	// Account Status & Usage
	SubscriptionTier string    `json:"subscriptionTier"`  // سطح اشتراک کاربر (مثلاً "free", "premium", "pro")
	AccountStatus    string    `json:"accountStatus"`     // وضعیت حساب کاربری: "active", "suspended", "deactivated"
	IsTwoFactorEnabled bool    `json:"isTwoFactorEnabled"`// آیا احراز هویت دو مرحله‌ای فعال است؟

	// Timestamps & Tracking
	LastActivity     time.Time `json:"lastActivity"`      // آخرین فعالیت عمومی کاربر در برنامه (برایengagement tracking)
	CreationDate     time.Time `json:"creationDate"`      // تاریخ ایجاد حساب کاربری
	ModificationDate time.Time `json:"modificationDate"`  // تاریخ آخرین تغییر در اطلاعات حساب کاربری

	// Generic/Extensible Metadata (for highly specific or future data)
	Metadata         map[string]string `json:"metadata"` // فیلد انعطاف‌پذیر برای ذخیره داده‌های اضافی یا خاص هر API/برنامه
}